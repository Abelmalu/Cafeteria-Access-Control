package models

import (
	"strings"
)

type Cafeteria struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Location string `json:"location" db:"cafeteria_location"`
}

func (c *Cafeteria) Validate() []string {
	var errs []string

	c.Name = strings.TrimSpace(c.Name)
	c.Location = strings.TrimSpace(c.Location)

	// Name validation
	if c.Name == "" {
		errs = append(errs, "cafeteria name is required")
	} else if len(c.Name) < 2 {
		errs = append(errs, "cafeteria name must be at least 2 characters")
	}

	if c.Location == "" {
		errs = append(errs, "cafeteria location is required")
	} else if len(c.Location) < 3 {
		errs = append(errs, "cafeteria location must be at least 3 characters")
	}

	if len(errs) > 0 {

		return errs
	}

	return nil
}
