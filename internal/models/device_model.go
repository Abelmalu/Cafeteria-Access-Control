package models

import "strings"

type Device struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	SerialNumber string `json:"serial_number" db:"serial_number"`
}

func (d *Device) Validate() []string {
	var errs []string
	const minNameLength = 2
	const maxNameLength = 100
	const minSerialLength = 5

	// --- 1. Name Validation ---
	d.Name = strings.TrimSpace(d.Name)
	if d.Name == "" {
		errs = append(errs, "Device Name is required.")
	} else {
		nameLength := len(d.Name)
		if nameLength < minNameLength {
			errs = append(errs, "Device Name must be at least 2 characters long.")
		}
		if nameLength > maxNameLength {
			errs = append(errs, "Device Name cannot exceed 100 characters.")
		}
	}

	// --- 2. Serial Number Validation ---
	d.SerialNumber = strings.TrimSpace(d.SerialNumber)
	if d.SerialNumber == "" {
		errs = append(errs, "Serial Number is required.")
	} else {
		// Minimum length check
		if len(d.SerialNumber) < minSerialLength {
			errs = append(errs, "Serial Number must be at least 5 characters long.")
		}

	}

	return errs
}
