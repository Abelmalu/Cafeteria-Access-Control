package mysql

import (
	"context"
	"database/sql"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type MySqlRepository struct {
	DB *sql.DB
}

// NewMySqlRepository creates a new MySqlRepository instance.
func NewMySqlRepository(db *sql.DB) *MySqlRepository {
	return &MySqlRepository{DB: db}
}

// This method implements the core.AccessRepository contract.
func (r *MySqlRepository) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	// This is where the actual SQL INSERT statement goes.
	query := `INSERT INTO students (student_id, rfid_tag, batch_id) VALUES ($1, $2, $3) RETURNING id;`
	// ... execute query using r.DB
	r.DB.Exec(query)
	return student, nil // return the created student
}
