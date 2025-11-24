package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type MySqlRepository struct {
	DB *sql.DB
}

// NewMySqlRepository creates a new MySqlRepository instance.
func NewMySqlRepository(db *sql.DB) *MySqlRepository {
	return &MySqlRepository{DB: db}
}

func (r *MySqlRepository) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	// 1. Removed RETURNING clause.
	// 2. Used MySQL-compatible '?' placeholders.
	// 3. Mapped all 7 fields correctly (removed the double 'rfid_tag').
	query := `
		INSERT INTO students (
			 first_name, middle_name, last_name, batch_id, rfid_tag, image_url
		) 
		VALUES ( ?, ?, ?, ?, ?, ?)`

	// Execute the query using ExecContext

	_, err := r.DB.ExecContext(ctx, query,

		student.FirstName,  // ?2
		student.MiddleName, // ?3
		student.LastName,   // ?4
		student.BatchId,    // ?5
		student.RFIDTag,    // ?6
		student.ImageURL,   // ?7
	)

	if err != nil {
		// Return an informative error if the insertion fails
		return nil, fmt.Errorf("mysql insert failed for student %s: %w", student.IdCard, err)
	}

	// Since we inserted all necessary data and don't need to fetch a new ID,
	// we just return the successfully inserted student object.
	return student, nil
}

func (r *MySqlRepository) CreateCafeteria(ctx context.Context, cafeteria *models.Cafeteria) (*models.Cafeteria, error) {

	query := `INSERT INTO cafeterias (id, name,careteria_location) VALUES (?,?,?)`

	_, err := r.DB.Exec(query,
		cafeteria.Id,       //?1
		cafeteria.Name,     //?2
		cafeteria.Location, //?3

	)

	if err != nil {

		return nil, err
	}

	return cafeteria, nil

}

func (r *MySqlRepository) CreateBatch(ctx context.Context, batch *models.Batch) (*models.Batch, error) {

	query := `INSERT INTO batches (name,cafeteria_id)  VALUES (?,?) `

	_, err := r.DB.Exec(query,

		batch.Name,
		batch.Cafeteria_id,
	)

	if err != nil {

		return nil, err
	}

	return batch, nil
}

func (r *MySqlRepository) CreateMeal(ctx context.Context, meal *models.Meal) (*models.Meal, error) {

	query := `INSERT INTO meals (name,start_time,end_time)  VALUES (?,?,?)`

	_, err := r.DB.Exec(query,
		meal.Name,
		meal.StartTime,
		meal.EndTime,
	)

	if err != nil {

		fmt.Println("create meal repo")
		return nil, err
	}
	fmt.Println("create meal repo not error")

	return meal, nil

}
func (r *MySqlRepository) RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error) {

	query := `INSERT INTO devices(name,serial_number) VALUES (?,?)`

	_, err := r.DB.Exec(query,
		device.Name,
		device.SerialNumber,
	)
	if err != nil {

		return nil, err
	}

	return device, nil
}

func (r *MySqlRepository) AttemptAccess(rfidTag string) (*models.Student, *models.Batch, error) {
	var student models.Student
	var batch models.Batch
	query := `SELECT * FROM students WHERE rfid_tag =?`

	studentRow := r.DB.QueryRow(query,
		rfidTag,
	)
	err := studentRow.Scan(
		&student.IdCard,
		&student.FirstName,
		&student.MiddleName,
		&student.LastName,
		&student.RFIDTag,
		&student.ImageURL,
		&student.BatchId,
	)

	if err != nil {

		return nil, &batch, errors.New("student not found")
	}

	//querying the database to get the
	batchQuery := `SELECT * FROM batches WHERE id=?`

	BatchRow := r.DB.QueryRow(
		batchQuery, student.BatchId,
	)
	BatchRowError := BatchRow.Scan(
		&batch.Id,
		&batch.Name,
		&batch.Cafeteria_id,
	)

	if BatchRowError != nil {

		return &student, nil, errors.New("not active batch")
	}

	return &student, &batch, nil

}

func (r *MySqlRepository) GetMeals() ([]models.Meal, error) {

	var meals []models.Meal

	mealQuery := `SELECT * FROM meals`

	mealRows, mealRowsErr := r.DB.Query(mealQuery)
	if mealRowsErr != nil {

		return nil, mealRowsErr
	}

	for mealRows.Next() {
		var m models.Meal
		// 3. Scan each row into a Meal struct
		if err := mealRows.Scan(&m.Id, &m.Name, &m.StartTime, &m.EndTime); err != nil {

			return nil, err
		}
		meals = append(meals, m)
	}
	if err := mealRows.Err(); err != nil {
		return nil, err
	}

	return meals, nil

}

func (r *MySqlRepository) GrantOrDenyAccess(currentDate string, studentId int, mealId int, cafeteriaId int) (string, error) {
	var mealLog models.MealAccessLog

	query := `SELECT * FROM meal_access_logs WHERE meal_id = ? AND student_id = ? AND scan_time = ?`

	result := r.DB.QueryRow(query,
		mealId,
		studentId,
		currentDate,
	)
	err := result.Scan(&mealLog.ID, &mealLog.ScanTime, &mealLog.Status, &mealLog.StudentID, &mealLog.CafeteriaID, &mealLog.MealID, &mealLog.DeviceID)

	// if no meal access log found insert the log
	if err != nil {
		insertQuery := `INSERT INTO meal_access_logs(scan_time,status,student_id,cafeteria_id,meal_id,device_id) VALUES(?,?,?,?,?,?)`
		fmt.Println(studentId)
		_, insertMealLogError := r.DB.Exec(insertQuery,
			currentDate,
			"Granted",
			studentId,
			cafeteriaId,
			mealId,
			1,
		)
		if insertMealLogError != nil {

			return "", insertMealLogError

		}
		return "Granted", nil

	}

	insertQuery := `INSERT INTO meal_access_logs(scan_time,status,student_id,cafeteria_id,meal_id,device_id) VALUES(?,?,?,?,?,?)`
	fmt.Println(studentId)
	_, insertMealLogError := r.DB.Exec(insertQuery,
		currentDate,
		"Denied",
		studentId,
		cafeteriaId,
		mealId,
		1,
	)
	if insertMealLogError != nil {

		return "couldn't save to meal logs", insertMealLogError

	}

	return "Denied", nil

}

func (r *MySqlRepository) GetCafeterias() ([]models.Cafeteria, error) {
	var cafeterias []models.Cafeteria

	query := `SELECT * FROM cafeterias`

	results, err := r.DB.Query(query)

	if err != nil {

		return nil, err
	}
	for results.Next() {
		var cafeteria models.Cafeteria

		results.Scan(&cafeteria.Id, &cafeteria.Name, &cafeteria.Location)

		cafeterias = append(cafeterias, cafeteria)

	}

	return cafeterias, nil

}

func (r *MySqlRepository) VerifyDevice(SerialNumber string) bool {

	var device models.Device

	query := `SELECT * FROM devices WHERE serial_number=?`

	result := r.DB.QueryRow(query, SerialNumber)
	err := result.Scan(&device.Id, &device.Name, &device.SerialNumber)

	if err == nil {

		return true
	}

	return false

}
