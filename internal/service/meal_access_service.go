package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type MealAccessService struct {
	repo core.MealAccessServiceRepository
}

func NewMealAccessService(repo core.MealAccessServiceRepository) *MealAccessService {

	return &MealAccessService{repo: repo}

}

func (ms *MealAccessService) AttemptAccess(rfidTag string, cafeteriaId string) (*models.Student, error) {

	if rfidTag == "" {

		return nil, errors.New("RFIDTag value empty")

	}
	if cafeteriaId == "" {

		return nil, errors.New("cafeteria id of the device is empty")

	}
	student, batch, err := ms.repo.AttemptAccess(rfidTag)

	if err != nil {

		return nil, err
	}

	deviceCafeteriaId, _ := strconv.Atoi(cafeteriaId)

	if deviceCafeteriaId == batch.Cafeteria_id {

		currentTime := time.Now()
		currentTimeFormatted := currentTime.Format("15:04:00")
		fmt.Println(currentTimeFormatted)

		meals, mealsErr := ms.repo.GetMeals()
		if mealsErr != nil {
			return student, mealsErr
		}
		var mealTime bool = false
		for _, value := range meals {

			startTime, _ := time.Parse("15:04:00", value.StartTime)

			finalStartTime := time.Date(
				currentTime.Year(),
				currentTime.Month(),
				currentTime.Day(),
				startTime.Hour(),
				startTime.Minute(),
				startTime.Second(),
				0,
				currentTime.Location())

			endTime, _ := time.Parse("15:04:00", value.EndTime)
			finalEndTime := time.Date(
				currentTime.Year(),
				currentTime.Month(),
				currentTime.Day(),
				endTime.Hour(),
				endTime.Minute(),
				endTime.Second(),
				0,
				currentTime.Location())

			if (currentTime.After(finalStartTime) || currentTime.Equal(finalStartTime)) &&
				(currentTime.Before(finalEndTime) || currentTime.Equal(finalEndTime)) {

				// 3. Found a match! Set true and BREAK the loop immediately.
				mealTime = true
				break
			}

		}
		if !mealTime {

			return student, errors.New("Not Meal Time")
		}

		return student, nil
	} else {

		return student, errors.New("Access Denied: Wrong Cafeteria.")
	}

}

// this method checks if the student can eat in the cafeteria
func CheckValidCafeteria(studentBatchCafeteria, deviceCafeteria string) (bool, error) {
	panic("unimplemented")
}

// checks if the current time is a meal time(breakfast,lunch,dinner)
func CheckMealTime(currentTime string) (*models.Meal, error) {
	panic("unimplemented")
}

// Grants or denies access to cafeteria for given student
func GrantOrDenyAccess(currentDate string, student *models.Student, mealId string, deviceId int) (bool, error) {
	panic("unimplemented")
}

// gets accesss logs
func GetAccessLog(date string) (*models.MealAccessLog, error) {
	panic("unimplemented")
}
