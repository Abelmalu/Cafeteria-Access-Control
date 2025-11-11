package models

type Batch struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Cafeteria_id int    `json:"cafeteria_id" db:"cafeteria_id"`
}
