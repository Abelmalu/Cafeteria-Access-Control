package mysql

import (
	"context"
	"database/sql"
	"fmt"

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
// that represents your database connection (*sql.DB or similar).

func (r *MySqlRepository) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	// 1. Removed RETURNING clause.
	// 2. Used MySQL-compatible '?' placeholders.
	// 3. Mapped all 7 fields correctly (removed the double 'rfid_tag').
	query := `
		INSERT INTO students (
			id_card, first_name, middle_name, last_name, batch_id, rfid_tag, image_url
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	// Execute the query using ExecContext
	_, err := r.DB.ExecContext(ctx, query,
		student.IdCard,     // ?1
		student.FirstName,  // ?2
		student.MiddleName, // ?3
		student.LastName,   // ?4
		student.BatchId,    // ?5
		student.RFIDTag,    // ?6
		student.ImageURL,   // ?7
	)

	if err != nil {
		// Return an informative error if the insertion fails
		return nil, fmt.Errorf("mysql insert failed for student %s: %w", student.IdCard, err)
	}

	// Since we inserted all necessary data and don't need to fetch a new ID,
	// we just return the successfully inserted student object.
	return student, nil
}
