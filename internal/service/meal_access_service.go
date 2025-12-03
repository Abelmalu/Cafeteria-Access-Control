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

func (ms *MealAccessService) AttemptAccess(rfidTag string, cafeteriaId string) (*models.Student, string, string, error) {

	if rfidTag == "" {

		return nil, "", "", errors.New("RFIDTag value empty")

	}
	if cafeteriaId == "" {

		return nil, "", "", errors.New("cafeteria id of the device is empty")

	}

	//get the student with the rfid tag and the batch
	student, batch, err := ms.repo.AttemptAccess(rfidTag)

	if err != nil {

		return nil, "", "", err
	}

	deviceCafeteriaId, _ := strconv.Atoi(cafeteriaId)

	//check if student is in the correct cafeteria
	if deviceCafeteriaId == batch.Cafeteria_id {

		currentTime := time.Now()

		//Getting meals to check if it is meal time
		meals, mealsErr := ms.repo.GetMeals()
		if mealsErr != nil {
			return student, "", batch.Name, mealsErr
		}
		var mealTime bool = false
		var mealID int
		for _, value := range meals {

			//change the meal.StartTime to time.Time
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

			//change the meal.EndTime to time.Time
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

			//checking if the current time blongs to a meal time window
			if (currentTime.After(finalStartTime) || currentTime.Equal(finalStartTime)) &&
				(currentTime.Before(finalEndTime) || currentTime.Equal(finalEndTime)) {

				// Found a match! Set true and BREAK the loop immediately.
				mealTime = true
				mealID = value.Id
				break
			}

		}
		// if not meal time return the student batch name and not meal time message
		if !mealTime {

			return student, "Not Meal Time", batch.Name, nil
		}
		currentDate := currentTime.Format("2006-01-02")

		// a method to grant or deny a student to the cafeteria
		grantReturn, grantError := ms.repo.GrantOrDenyAccess(currentDate, student.IdCard, mealID, deviceCafeteriaId)

		if grantError != nil {

			return student, "", batch.Name, grantError
		}

		fmt.Println(grantReturn)

		return student, grantReturn, batch.Name, nil
	} else {

		return student, "Wrong Cafeteria", batch.Name, nil
	}

}

func (ms *MealAccessService) GetCafeterias() ([]models.Cafeteria, error) {

	cafeterias, err := ms.repo.GetCafeterias()

	if err != nil {

		return nil, err
	}

	return cafeterias, nil
}

func (ms *MealAccessService) VerifyDevice(SerialNumber string) bool {

	if SerialNumber == "" {

		return false

	}

	exists := ms.repo.VerifyDevice(SerialNumber)

	if exists {

		return true
	}

	return false

}
