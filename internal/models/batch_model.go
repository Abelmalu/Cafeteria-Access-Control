package models

import (
	"strings"
)

type Batch struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Cafeteria_id int    `json:"cafeteria_id" db:"cafeteria_id"`
}

func (b *Batch) Validate() []string {

	var errs []string

	b.Name = strings.TrimSpace(b.Name)

	if b.Name == "" {
		errs = append(errs, "Batch name is required")

	} else if len(b.Name) < 2 {

		errs = append(errs, "Batch name must be at least two characters")

	}

	if b.Cafeteria_id <= 0 {

		errs = append(errs, "cafeteria ID must be greater than zero")

	}

	if len(errs) > 0 {

		return errs
	}

	return nil

}
