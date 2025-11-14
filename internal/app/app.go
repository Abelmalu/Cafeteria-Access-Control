package app

import (
	"database/sql"
	"fmt"
	"github.com/abelmalu/CafeteriaAccessControl/config"
	"github.com/abelmalu/CafeteriaAccessControl/internal/api"
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/repository/mysql"
	"github.com/abelmalu/CafeteriaAccessControl/internal/repository/postgres"
	"github.com/abelmalu/CafeteriaAccessControl/internal/service"

	"github.com/go-chi/chi/v5"
	mysqlDriver "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

// App holds the application dependencies and router.
// This is the main state of the running application.
type App struct {
	Config *config.Config
	Router *chi.Mux // The core Go HTTP router
	DB     *sql.DB  // The database connection pool
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

	if DBerr != nil {

		fmt.Println(DBerr)
	}

	app := &App{Config: config, Router: chi.NewRouter(), DB: currentDBConnection}

	// 2. Setup all layers and routes: Performs the dependency injection.
	app.setupRoutes()

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

	repo, err := NewRepositoryFactory(a.Config.DBType, a.DB)
	if err != nil {
		log.Fatalf("FATAL: Failed to initialize repository for type %s: %v", a.Config.DBType, err)
	}
	log.Println("INFO: Abstract Repository initialized with concrete implementation:", a.Config.DBType)

	// Service initialization creates the 'adminSvc' variable
	adminSvc := service.NewAdminService(repo)

	// ✅ FIX 1: Pass the initialized service variable (adminSvc) into the handler constructor
	adminHandler := api.NewAdminHandler(adminSvc)

	// ✅ FIX 2: Call the method on the initialized handler variable (adminHandler)
	a.Router.Post("/api/admin/create/student", http.HandlerFunc(adminHandler.CreateStudent))
}

// Run starts the HTTP server on the configured port.
func (a *App) Run() {
	log.Printf("Server listening on :%s", a.Config.ServerPort)
	ServerPort := strconv.Itoa(a.Config.ServerPort)

	// The router (a.Router) handles all the routes and middleware defined above.
	if err := http.ListenAndServe(":"+ServerPort, a.Router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// instantiates the correct concrete implementation based on the dbType string.
func NewRepositoryFactory(dbType string, db *sql.DB) (core.AccessRepository, error) {
	switch dbType {
	case "mysql":
		// Returns the concrete *mysql.MySqlRepository, which implements core.Repository
		return mysql.NewMySqlRepository(db), nil
	case "postgres":
		// Returns the concrete *postgres.PostgresRepository, which implements core.Repository
		return postgres.NewPostgresRepository(db), nil
	// You can add 'inmemory' for testing or 'sqlite' here
	default:
		return nil, fmt.Errorf("unsupported database repository type specified: %s", dbType)
	}
}
