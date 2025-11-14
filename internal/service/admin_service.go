package service

import (
	"context"
	"errors"
	//"time"

	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type AdminService struct {
	repo core.AccessRepository
}

// NewAdminService creates a new instance of AdminService.
func NewAdminService(repo core.AccessRepository) *AdminService {
	return &AdminService{repo: repo}
}

// CreateBatch implements core.AdminService.
func (s *AdminService) CreateBatch(ctx context.Context, student *models.Student) (*models.Student, error) {
	panic("unimplemented")
}

// CreateCafeteria implements core.AdminService.
func (s *AdminService) CreateCafeteria(ctx context.Context, location *models.Device) (*models.Cafeteria, error) {
	panic("unimplemented")
}

// CreateMeal implements core.AdminService.
func (s *AdminService) CreateMeal(ctx context.Context, student *models.Student) (*models.Student, error) {
	panic("unimplemented")
}

// RegisterDevice implements core.AdminService.
func (s *AdminService) RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error) {
	panic("unimplemented")
}

func (s *AdminService) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	// --- Service Layer Validation ---
	if student.IdCard == "" || student.BatchId == "" || student.RFIDTag == "" {
		return nil, errors.New("admin: student ID, batch ID, and RFID tag are required")
	}

	// NOTE: In a complete system, we would check for duplicate RFIDTag or StudentID
	// before calling the repository's CreateStudent method to avoid SQL errors.

	// The repository handles the actual database INSERT.
	return s.repo.CreateStudent(ctx, student)
}

// func CreateBatch(ctx context.Context, student *models.Student) (*models.Student, error)       {}
// func RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error)       {}
// func CreateMeal(ctx context.Context, student *models.Student) (*models.Student, error)        {}
// func CreateCafeteria(ctx context.Context, location *models.Device) (*models.Cafeteria, error) {}
