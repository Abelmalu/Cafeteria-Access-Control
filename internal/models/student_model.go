package models

import (
	"fmt"
	"strings"
)

type Student struct {
	IdCard     int    `json:"id_card" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	LastName   string `json:"last_name" db:"last_name"`
	BatchId    int    `json:"batch_id" db:"batch_id"`
	RFIDTag    string `json:"rfid_tag" db:"rfid_tag"`
	ImageURL   string `json:"image_url" db:"image_url"` // path or URL to the image
}

func (s *Student) Validate() []string {
	var errs []string

	// --- 1. ID and Foreign Key Validation ---

	// IdCard: Should not be zero if being updated/read (though often database-managed)
	// We'll focus on the required foreign key (BatchId) for insertion validation.
	if s.BatchId <= 0 {
		errs = append(errs, "Batch ID is required and must be a positive integer.")
	}

	// --- 2. String Field Validation (Required Fields) ---

	// First Name
	s.FirstName = strings.TrimSpace(s.FirstName)
	if s.FirstName == "" {
		errs = append(errs, "First Name is required.")
	}

	// Last Name
	s.LastName = strings.TrimSpace(s.LastName)
	if s.LastName == "" {
		errs = append(errs, "Last Name is required.")
	}

	// RFID Taghhh
	s.RFIDTag = strings.TrimSpace(s.RFIDTag)
	if s.RFIDTag == "" {
		errs = append(errs, "RFID Tag is required.")
	} else if len(s.RFIDTag) < 8 || len(s.RFIDTag) > 32 {
		// Example length check for common RFID tags
		errs = append(errs, fmt.Sprintf("RFID Tag must be between 8 and 32 characters (currently %d).", len(s.RFIDTag)))
	}

	// --- 3. Optional Fields Validation ---

	// Middle Name is often optional, but we can clean it up.
	s.MiddleName = strings.TrimSpace(s.MiddleName)

	// Image URL (If the URL is provided, perform basic formatting checks)
	s.ImageURL = strings.TrimSpace(s.ImageURL)
	if s.ImageURL != "" && !strings.HasPrefix(s.ImageURL, "http") && !strings.HasPrefix(s.ImageURL, "/") {
		errs = append(errs, "Image URL must be a valid path (starting with 'http' or '/').")
	}

	return errs
}
