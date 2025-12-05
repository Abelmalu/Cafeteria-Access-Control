package api_test

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/abelmalu/CafeteriaAccessControl/internal/api"
// 	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
// )

// type MockMealAccessService struct {
// 	err          error
// 	student      models.Student
// 	cafeterias   []models.Cafeteria
// 	accessStatus string
// 	batchName    string
// }

// // AttemptAccess implements core.MealAccessServiceRepository.
// func (m MockMealAccessService) AttemptAccess(rfidTag string, cafeteriaId string) (*models.Student, string, string, error) {
// 	panic("")
// }

// // GetMeals implements core.MealAccessServiceRepository.
// func (m MockMealAccessService) GetMeals() ([]models.Meal, error) {
// 	panic("")
// }

// // GrantOrDenyAccess implements core.MealAccessServiceRepository.
// func (m MockMealAccessService) GrantOrDenyAccess(currentDate string, studentId int, mealId int, cafeteriaId int) (string, error) {
// 	panic("")

// }

// func (m MockMealAccessService) GetCafeterias() ([]models.Cafeteria, error) {
// 	return m.cafeterias, m.err
// }

// func (r *MockMealAccessService) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

// 	panic("")
// }

// func (r *MockMealAccessService) CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error) {

// 	panic("")

// }

// func (r *MockMealAccessService) CreateBatch(ctx context.Context, batch *models.Batch) (*models.Batch, error) {

// 	panic("")
// }

// func (r *MockMealAccessService) CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error) {
// 	panic("unimplemented")

// }

// // VerifyDevice implements core.MealAccessServiceRepository.
// func (m MockMealAccessService) VerifyDevice(SerialNumber string) bool {

// 	panic("")
// }

// /*

// 	Testing begins here

// */

// func TestMealAccessHandler_GetCafeterias_Success(t *testing.T) {
// 	mockSvc := MockMealAccessService{

// 		cafeterias: []models.Cafeteria{
// 			{Id: 1, Name: "Main Hall", Location: "Building A"},
// 			{Id: 2, Name: "North Cafe", Location: "Building B"},
// 		},
// 		err: nil,
// 	}
// 	mealAccessHandler := api.NewMealAccessHandler(mockSvc)

// 	req := httptest.NewRequest(http.MethodGet, "/api/cafeterias", nil)
// 	recorder := httptest.NewRecorder()
// 	mealAccessHandler.GetCafeterias(recorder, req)

// 	if recorder.Code != http.StatusAccepted {
// 		t.Errorf("expected status %d, got %d", http.StatusAccepted, recorder.Code)
// 	}

// 	var response []models.Cafeteria
// 	err := json.Unmarshal(recorder.Body.Bytes(), &response)
// 	if err != nil {
// 		t.Fatalf("failed to decode json: %v", err)
// 	}

// 	// 8. Validate response content
// 	if len(response) != 2 {
// 		t.Errorf("expected 2 cafeterias, got %d", len(response))
// 	}

// 	if response[0].Name != "Main Hall" {
// 		t.Errorf("expected Cafe A, got %s", response[0].Name)
// 	}

// }

// func TestMealAccessHandler_GetCafeterias_Error(t *testing.T) {

// 	mockScv := MockMealAccessService{
// 		err: errors.New("Couldn't fetch Cafeterias"),
// 	}

// 	mealAccessHandler := api.NewMealAccessHandler(mockScv)

// 	request := httptest.NewRequest(http.MethodGet, "/api/cafeterias", nil)
// 	recorder := httptest.NewRecorder()

// 	mealAccessHandler.GetCafeterias(recorder, request)

// 	if recorder.Code != http.StatusInternalServerError {

// 		t.Fatalf("unexpected status code %v", recorder.Code)
// 	}

// 	expected := `{"status":"error","message":"Couldn't fetch Cafeterias"}`
// 	if strings.TrimSpace(recorder.Body.String()) != expected {
// 		t.Errorf("expected %s, got %s", expected, recorder.Body.String())
// 	}

// }
