package postgres

import (
	"context"
	"database/sql"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type PostgresRepository struct {
	DB *sql.DB
}

// NewMySqlRepository creates a new MySqlRepository instance.
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

// This method implements the core.AccessRepository contract.
func (r *PostgresRepository) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	// This is where the actual SQL INSERT statement goes.
	query := `INSERT INTO students (student_id, rfid_tag, batch_id) VALUES ($1, $2, $3) RETURNING id;`
	// ... execute query using r.DB
	r.DB.Exec(query)
	return student, nil // return the created student
}

func (r *PostgresRepository) CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error) {

	panic("unimplemented")
}


func (r *PostgresRepository) CreateBatch(ctx context.Context, student *models.Batch) (*models.Batch, error){

	panic("unimplemented")
}

func (r *PostgresRepository) CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error){


	panic("unimplemented")
	
}

func (r *PostgresRepository) RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error){

	panic("unimplemented")
}


func (r *PostgresRepository) GetMeals() ([]models.Meal, error) {

	panic("unimplemented")
}






func (r *PostgresRepository) AttemptAccess(rfidTag string) (*models.Student,*models.Batch, error) {
	panic("unimplemented")

}

