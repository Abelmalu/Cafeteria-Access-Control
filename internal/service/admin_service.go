package service

import (
	"context"
	"errors"
	"time"

	//"time"

	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type AdminService struct {
	repo core.AdminRepository
}

// NewAdminService creates a new instance of AdminService.
func NewAdminService(repo core.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}

// CreateBatch implements core.AdminService.
func (s *AdminService) CreateBatch(ctx context.Context, batch *models.Batch) (*models.Batch, error) {
	if batch.Cafeteria_id == 0 || batch.Name == "" {

		return nil, errors.New("cafeteria_id and/or batch name are required ")
	}

	_, err := s.repo.CreateBatch(ctx, batch)
	if err != nil {

		return nil, err
	}

	return batch, nil

}

// CreateCafeteria implements core.AdminService.
func (s *AdminService) CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error) {
	if cafeteria.Name == "" || cafeteria.Location == "" {

		return nil, errors.New("Cafeteria id and/or Cafeteria name are required")

	}
	_, err := s.repo.CreateCafeteria(ctx, cafeteria)
	if err != nil {

		return nil, err
	}

	return cafeteria, nil
}

// CreateMeal implements core.AdminService.
func (s *AdminService) CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error) {
	if meal.Name == "" || meal.StartTime == "" || meal.EndTime == "" {

		return nil, errors.New("meal Name/StartTime/EndTime is required")

	}

	// parsing and formatting for start time
	startTime, startTimeParseErr := time.Parse("3:04 PM", meal.StartTime)
	if startTimeParseErr != nil {

		return nil, errors.New("invalid start time")
	}
	meal.StartTime = startTime.Format("15:04:05")

	// parsing and formatting for end time
	endTime, endTimeParseErr := time.Parse("3:04 PM", meal.EndTime)
	if endTimeParseErr != nil {

		return nil, errors.New("invalid end time")
	}
	meal.EndTime = endTime.Format("15:04:05")

	_, err := s.repo.CreateMeal(ctx, meal)

	if err != nil {

		return nil, err
	}

	return meal, nil

}

// RegisterDevice implements core.AdminService.
func (s *AdminService) RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error) {
	if device.Name == "" || device.SerialNumber == "" {

		return nil, errors.New("device name and serialnumber can not be nul and less than zero")
	}

	deviceReturned, err := s.repo.RegisterDevice(ctx, device)

	if err != nil {

		return nil, err
	}

	return deviceReturned, nil

}

func (s *AdminService) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	// --- Service Layer Validation ---
	if student.FirstName == "" || student.MiddleName == "" || student.LastName == "" || student.BatchId == 0 || student.RFIDTag == "" {
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
