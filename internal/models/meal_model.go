package models

// struct defining meals and their time

type Meal struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime   string `json:"end_time" db:"end_time"`
}
