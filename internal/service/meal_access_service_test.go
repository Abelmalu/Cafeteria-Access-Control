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
	cafeterias   []models.Cafeteria
	student      models.Student
	batch        models.Batch
	meals        []models.Meal
	grantMessage string
	exists       bool
	err          error
}

// AttemptAccess implements core.MealAccessServiceRepository.
func (m MockMealAccessRepo) AttemptAccess(rfidTag string) (*models.Student, *models.Batch, error) {
	return &m.student, &m.batch, m.err
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

	return m.exists

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

func TestMealAccessService_VerifyDevice_Success(t *testing.T) {

	mockRepo := MockMealAccessRepo{

		exists: true,
	}

	svc := service.NewMealAccessService(mockRepo)

	exists := svc.VerifyDevice("SerialNO2")

	if exists == false {

		t.Fatalf("unexpected error expecting true got %v ", exists)
	}

}

func TestMealAccessService_VerifyDevices_Error(t *testing.T) {

	mockRepo := MockMealAccessRepo{

		exists: false,
	}

	svc := service.NewMealAccessService(mockRepo)

	exists := svc.VerifyDevice("SerialNOError")

	if exists == true {

		t.Fatalf("expecting false got %v", exists)
	}

}

func TestMealAccessService_VerifyDevice_EmptySerialNumber(t *testing.T) {
	mockRepo := MockMealAccessRepo{
		exists: false,
	}

	svc := service.NewMealAccessService(mockRepo)

	// Test with empty serial number
	exists := svc.VerifyDevice("")

	if exists == true {
		t.Fatal("expected false for empty serial number, got true")
	}
}

// func TestMealAccessService_AttemptAccess_EmptyRFIDTag(t *testing.T) {
// 	mockRepo := MockMealAccessRepo{

// 		meals: []models.Meal{
// 			{
// 				Name:      "kurs",
// 				StartTime: "7:00:00",
// 				EndTime:   "2:00:00",
// 			},
// 			{
// 				Name:      "mesa",
// 				StartTime: "8:00:00",
// 				EndTime:   "9:00:00",
// 			},
// 		},
// 		student: models.Student{

// 			FirstName:  "abe",
// 			MiddleName: "gsa",
// 			LastName:   "hello",
// 			RFIDTag:    "fc:22",
// 			BatchId:    1,
// 			ImageURL:   "assets",
// 		},
// 		grantMessage: "Granted",
// 	}

// }
// func TestMealAccessService_AttemptAccess_EmptycafeteriaId(t *testing.T) {

// }

func TestMealAccessService_AttemptAccess_Success(t *testing.T) {

	mockRepo := MockMealAccessRepo{

		meals: []models.Meal{
			{
				Name:      "kurs",
				StartTime: "7:00:00",
				EndTime:   "17:00:00",
			},
			{
				Name:      "mesa",
				StartTime: "8:00:00",
				EndTime:   "9:00:00",
			},
		},
		student: models.Student{

			FirstName:  "abe",
			MiddleName: "gsa",
			LastName:   "hello",
			RFIDTag:    "fc:22",
			BatchId:    1,
			ImageURL:   "assets",
		},
		batch: models.Batch{
			Name:         "2025",
			Cafeteria_id: 1,
		},
		grantMessage: "Granted",
	}

	svc := service.NewMealAccessService(mockRepo)
	_, accessStatus, _, err := svc.AttemptAccess("fc:22", "1")

	if err != nil {

		t.Fatalf("unexpected error %v", err)
	}

	if !(accessStatus == "Granted") || !(accessStatus == "Denied") {

		t.Fatalf("unexpected error message should be Granted or Denied")

	}

}
