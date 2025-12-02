package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
	"github.com/abelmalu/CafeteriaAccessControl/internal/service"
)

// Mock Repository
type MockMealAccessRepo struct {
	cafeterias []models.Cafeteria
	err        error
}

// AttemptAccess implements core.MealAccessServiceRepository.
func (m MockMealAccessRepo) AttemptAccess(rfidTag string) (*models.Student, *models.Batch, error) {
	panic("unimplemented")
}

// GetMeals implements core.MealAccessServiceRepository.
func (m MockMealAccessRepo) GetMeals() ([]models.Meal, error) {
	panic("unimplemented")
}

// GrantOrDenyAccess implements core.MealAccessServiceRepository.
func (m MockMealAccessRepo) GrantOrDenyAccess(currentDate string, studentId int, mealId int, cafeteriaId int) (string, error) {
	panic("unimplemented")
}

func (m MockMealAccessRepo) GetCafeterias() ([]models.Cafeteria, error) {
	return m.cafeterias, m.err
}

func (r *MockMealAccessRepo) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

	panic("")
}

func (r *MockMealAccessRepo) CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error) {

	panic("")

}

func (r *MockMealAccessRepo) CreateBatch(ctx context.Context, batch *models.Batch) (*models.Batch, error) {

	panic("")
}

func (r *MockMealAccessRepo) CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error) {
	panic("")

}

// VerifyDevice implements core.MealAccessServiceRepository.
func (m MockMealAccessRepo) VerifyDevice(SerialNumber string) bool {
	panic("unimplemented")
}

//
// Tests
//

func TestMealAccessService_GetCafeterias_Success(t *testing.T) {

	mockRepo := MockMealAccessRepo{
		cafeterias: []models.Cafeteria{
			{Id: 1, Name: "Main Hall", Location: "Building A"},
			{Id: 2, Name: "North Cafe", Location: "Building B"},
		},
		err: nil,
	}

	svc := service.NewMealAccessService(mockRepo)

	result, err := svc.GetCafeterias()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 cafeterias, got %d", len(result))
	}

	if result[0].Name != "Main Hall" {
		t.Errorf("expected first cafeteria name 'Main Hall', got '%s'", result[0].Name)
	}
}

func TestMealAccessService_GetCafeterias_Error(t *testing.T) {

	mockRepo := MockMealAccessRepo{
		cafeterias: nil,
		err:        errors.New("database error"),
	}

	svc := service.NewMealAccessService(mockRepo)

	_, err := svc.GetCafeterias()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "database error" {
		t.Errorf("unexpected error message: %v", err)
	}
}
