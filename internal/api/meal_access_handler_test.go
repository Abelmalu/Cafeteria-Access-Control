package api_test

import (
	"context"
	"testing"

	"github.com/abelmalu/CafeteriaAccessControl/internal/api"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type MockMealAccessService struct {
	err          error
	student      models.Student
	cafeterias   []models.Cafeteria
	accessStatus string
	batchName    string
}

// AttemptAccess implements core.MealAccessServiceRepository.
func (m MockMealAccessService) AttemptAccess(rfidTag string, cafeteriaId string) (*models.Student, string, string, error) {
	panic("")
}

// GetMeals implements core.MealAccessServiceRepository.
func (m MockMealAccessService) GetMeals() ([]models.Meal, error) {
	panic("")
}

// GrantOrDenyAccess implements core.MealAccessServiceRepository.
func (m MockMealAccessService) GrantOrDenyAccess(currentDate string, studentId int, mealId int, cafeteriaId int) (string, error) {
	panic("")

}

func (m MockMealAccessService) GetCafeterias() ([]models.Cafeteria, error) {
	panic("")
}

func (r *MockMealAccessService) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

	panic("")
}

func (r *MockMealAccessService) CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error) {

	panic("")

}

func (r *MockMealAccessService) CreateBatch(ctx context.Context, batch *models.Batch) (*models.Batch, error) {

	panic("")
}

func (r *MockMealAccessService) CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error) {
	panic("unimplemented")

}

// VerifyDevice implements core.MealAccessServiceRepository.
func (m MockMealAccessService) VerifyDevice(SerialNumber string) bool {

	panic("")
}

/*

	Testing begins here


*/

func TestMealAccessHandler_GetCafeterias_Success(t *testing.T) {
	mockSvc := MockMealAccessService{

		cafeterias: []models.Cafeteria{
			{Id: 1, Name: "Main Hall", Location: "Building A"},
			{Id: 2, Name: "North Cafe", Location: "Building B"},
		},
		err: nil,
	}
	mealAccessHandler := api.NewMealAccessHandler(mockSvc)

	mealAccessHandler.GetCafeterias()

}
