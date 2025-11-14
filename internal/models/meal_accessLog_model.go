package models

import "time"

// AccessLog represents an audit record of a single attempt to access the cafeteria.
// This is the main transactional table used for both historical reporting and
// immediate denial logic (preventing duplicate meals).
type MealAccessLog struct {
	ID        int64     `json:"id" db:"id"`                   // Primary Key, auto-incremented
	StudentID string    `json:"student_id" db:"student_id"`   // Foreign Key to the Student table
	MealID    int       `json:"meal_id" db:"meal_id"`         // Foreign Key to the MealTime table (e.g., 1=Lunch)
	ScanTime  string    `json:"scan_time" db:"scan_time"`     // Exact time the card was read
	DeviceID  int       `json:"location_id" db:"location_id"` // Reader location (if multiple entry points)
	Status    string    `json:"status" db:"status"`           // The outcome: "SUCCESS" or "DENIED"
	CreatedAt time.Time `json:"created_at" db:"created_at"`   // Timestamp for record creation
}
