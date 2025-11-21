package models

type Device struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	SerialNumber string `json:"serial_number" db:"serial_number"`
}
