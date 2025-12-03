package core

import (
	"context"

	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type AdminRepository interface {
	// GetStudentByRfidTag(tag string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	CreateBatch(ctx context.Context, batch *models.Batch) (*models.Batch, error)
	CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error)
	RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error)
	CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error)
}
