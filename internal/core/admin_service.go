package core

import (
	"context"

	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

// AdminService defines the contract for administrative functions
// such as creating user accounts, registering hardware, and setting policies.
type AdminService interface {
	// RegisterDevice creates a new physical scanner device record.
	// It is responsible for internal logic like generating a secure API key before saving.
	// Returns the created Device model, which includes the generated API key.
	RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error)

	// CreateBatch adds a new batch identity to the database.
	CreateBatch(ctx context.Context, batch *models.Batch) (*models.Batch, error)

	// CreateStudent adds a new student identity to the database.
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)

	// CreateMeal adds a new meal identity to the database.
	CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error)

	// creates  cafeteria/dining hall record.
	CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error)
}
