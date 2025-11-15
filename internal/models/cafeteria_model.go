package models

type Cafeteria struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Location string `json:"location" db:"location"`
}
