package app

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/abelmalu/CafeteriaAccessControl/config"
	"github.com/abelmalu/CafeteriaAccessControl/internal/api"
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
	"github.com/abelmalu/CafeteriaAccessControl/internal/repository/mysql"
	"github.com/abelmalu/CafeteriaAccessControl/internal/repository/postgres"
	"github.com/abelmalu/CafeteriaAccessControl/internal/service"
	"github.com/brianvoe/gofakeit/v6"
	mysqlDriver "github.com/go-sql-driver/mysql"

	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5"
)

//go:embed sql/ddl.sql
var ddlFile string

var mysqlErr *mysqlDriver.MySQLError

//go:embed static
var embeddedStaticFS embed.FS

// App holds the application dependencies and router.
// This is the main state of the running application.
type App struct {
	Config   *config.Config
	StaticFS embed.FS
	Router   *chi.Mux // The core Go HTTP router
	DB       *sql.DB
	//MealAccessSvc *service.MealAccessService
}

// NewApp loads configuration and initializes the application structure.
// This function performs the core setup and dependency injection.
func NewApp() (*App, error) {

	config, err := config.LoadConfig()

	if err != nil {

		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// 1. Initialize Database Connection: Must be done first as all other layers depend on it.
	currentDBConnection, DBerr := initDB(config)

	migrationsError := runMigrations(currentDBConnection)

	if migrationsError != nil {

		log.Fatal(migrationsError)
	}

	if DBerr != nil {

		fmt.Println(DBerr)
	}

	router := chi.NewRouter()

	// --- 2. CONFIGURE AND USE CORS MIDDLEWARE ---
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: are the domains/IPs allowed to access your API.
		// Use "*" during development to allow all, then change to specific frontend domains in production.
		AllowedOrigins: []string{"*"},

		// AllowedMethods: specify the HTTP methods that are permitted.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},

		// AllowedHeaders: allow the frontend to send critical headers like Content-Type.
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},

		// Allow the browser to send cookies/auth headers if needed (for later).
		AllowCredentials: true,

		// MaxAge: how long the browser can cache the preflight response (in seconds).
		MaxAge: 300,
	}))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	app := &App{Config: config, Router: router, DB: currentDBConnection, StaticFS: embeddedStaticFS}
	fmt.Println(embeddedStaticFS)
	fmt.Println("printing  the embeded file static ")
	// 2. Setup all layers and routes: Performs the dependency injection.
	app.setupRoutes()
	app.CreateDummyStudents()

	return app, nil
}

// initDB establishes the database connection and ensures connection pool health.
func initDB(cfg *config.Config) (*sql.DB, error) {
	var connStr string
	var driverName string

	// 1. Determine the driver and connection string based on configuration
	switch cfg.DBType {
	case "mysql":
		driverName = "mysql"
		// Using the driver's native Config struct is robust
		mysqlCfg := mysqlDriver.Config{
			User:                 cfg.DBUser,
			Passwd:               cfg.DBPassword,
			Net:                  "tcp",
			Addr:                 cfg.DBHost,
			DBName:               cfg.DBName,
			AllowNativePasswords: true,
			ParseTime:            true,
			MultiStatements:      true,
		}
		connStr = mysqlCfg.FormatDSN()
	case "postgres":
		driverName = "pgx"
		// PostgreSQL connection string format (e.g., "host=... user=... password=...")
		connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.DBType)
	}

	// 2. Open the connection (standard library function)
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("opening database connection for %s: %w", driverName, err)
	}

	// 3. Apply pooling and verify connection (this part is vendor-agnostic)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5) // Reduced from 25 to 5 for slightly better resource usage when idle
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("pinging %s database: %w", driverName, err)
	}

	return db, nil
}

// setupRoutes initializes all concrete implementations and wires them together.
func (a *App) setupRoutes() {

	Adminrepo, MealAccessRepo, err := NewRepositoryFactory(a.Config.DBType, a.DB)
	if err != nil {
		log.Fatalf("FATAL: Failed to initialize repository for type %s: %v", a.Config.DBType, err)
	}
	log.Println("INFO: Abstract Repository initialized with concrete implementation:", a.Config.DBType)

	mealAccessSvc := service.NewMealAccessService(MealAccessRepo)
	mealAccessHandler := api.NewMealAccessHandler(mealAccessSvc)

	// Static file router
	staticSubFS, _ := fs.Sub(embeddedStaticFS, "static")
	fsHandler := http.FileServer(http.FS(staticSubFS))
	a.Router.Handle("/static/*", http.StripPrefix("/static/", fsHandler))
	uploadHandler := http.FileServer(http.Dir(os.Getenv("UPLOAD_DIR")))
	a.Router.Handle("/uploads/*", http.StripPrefix("/uploads/", uploadHandler))

	//meal Access routes starts here
	a.Router.Get("/api/mealaccess/{sutdentRfid}/{cafeteriaId}", http.HandlerFunc(mealAccessHandler.AttemptAccess))
	a.Router.Get("/api/cafeterias", http.HandlerFunc(mealAccessHandler.GetCafeterias))
	a.Router.Get("/api/device/verify/{SerialNumber}", http.HandlerFunc(mealAccessHandler.VerifyDevice))
	//meal access routes ends here

	// Service initialization creates the 'adminSvc' variable
	adminSvc := service.NewAdminService(Adminrepo)

	// ✅ FIX 1: Pass the initialized service variable (adminSvc) into the handler constructor
	adminHandler := api.NewAdminHandler(adminSvc)

	// admin routes
	a.Router.Post("/api/admin/create/cafeteria", http.HandlerFunc(adminHandler.CreateCafeteria))
	a.Router.Post("/api/admin/create/batch", http.HandlerFunc(adminHandler.CreateBatch))
	a.Router.Post("/api/admin/create/student", http.HandlerFunc(adminHandler.CreateStudent))
	a.Router.Post("/api/admin/create/meal", http.HandlerFunc(adminHandler.CreateMeal))
	a.Router.Post("/api/admin/register/device", http.HandlerFunc(adminHandler.RegisterDevice))
	// admin routes ends here

	// <-- FIX IS HERE

}

func (a *App) CreateDummyStudents() {

	fmt.Println(gofakeit.Name())
	var total int = 10000

	for i := 0; i <= total; i++ {
		student := models.Student{}

		student.FirstName = gofakeit.FirstName()
		student.MiddleName = gofakeit.MiddleName()
		student.LastName = gofakeit.LastName()
		student.BatchId = 1
		student.RFIDTag = fmt.Sprintf("79:22 %d", i)
		student.ImageURL = fmt.Sprintf(
			"https://picsum.photos/seed/%d/200/200",
			i,
		)

		query := `
		INSERT INTO students (
			 first_name, middle_name, last_name, batch_id, rfid_tag, image_url
		) 
		VALUES ( ?, ?, ?, ?, ?, ?)`

		// Execute the query using ExecContext

		_, err := a.DB.Exec(query,

			student.FirstName,  // ?2
			student.MiddleName, // ?3
			student.LastName,   // ?4
			student.BatchId,    // ?5
			student.RFIDTag,    // ?6
			student.ImageURL,   // ?7
		)

		if err != nil {
			// Return an informative error if the insertion fails
			if errors.As(err, &mysqlErr) {

				switch mysqlErr.Number {

				case 1062:
					fmt.Println("Student already exists with this rfid tag")
				default:
					fmt.Println(err)
				}

			}
		}

	}

}

// Run starts the HTTP server on the configured port.
func (a *App) Run() {
	log.Printf("Server listening on :%v", a.Config.ServerPort)
	ServerPort := strconv.Itoa(a.Config.ServerPort)

	// The router (a.Router) handles all the routes and middleware defined above.
	if err := http.ListenAndServe("192.168.100.169:"+ServerPort, a.Router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// instantiates the correct concrete implementation based on the dbType string.
func NewRepositoryFactory(dbType string, db *sql.DB) (core.AdminRepository, core.MealAccessServiceRepository, error) {
	switch dbType {
	case "mysql":
		// Returns the concrete *mysql.MySqlRepository, which implements core.Repository
		return mysql.NewMySqlRepository(db), mysql.NewMySqlRepository(db), nil
	case "postgres":
		// Returns the concrete *postgres.PostgresRepository, which implements core.Repository
		return postgres.NewPostgresRepository(db), postgres.NewPostgresRepository(db), nil
	// You can add 'inmemory' for testing or 'sqlite' here

	default:
		return nil, nil, fmt.Errorf("unsupported database repository type specified: %s", dbType)
	}
}

// --- Migration Helpers ---

// runMigrations reads the content of the DDL file and executes it.
// This function ensures the base schema is present before the app starts.
// runMigrations executes the DDL script that was embedded into the 'ddlContent' string.
func runMigrations(db *sql.DB) error {
	log.Printf("INFO: Running migrations from embedded DDL content...")

	if ddlFile == "" {
		return fmt.Errorf("ddl content is empty; check the //go:embed directive and file content")
	}

	// Execute the entire content of the SQL file as a single block.
	// We execute the string content directly, avoiding os.Open.
	_, err := db.Exec(ddlFile)
	if err != nil {
		// Log the error returned by the database driver
		return fmt.Errorf("failed to execute DDL: %w", err)
	}

	return nil
}
