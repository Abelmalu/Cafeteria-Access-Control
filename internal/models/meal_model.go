package models

import (
	"fmt"
	"time"
)

// struct defining meals and their time

type Meal struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime   string `json:"end_time" db:"end_time"`
}

// check if current time is in the Meal's start/end window.

func (m *Meal) IsIt(t time.Time) (bool, error) {
	// Parse the time components from the database strings to Time type
	// time.TimeOnly is formatted like "1:1:00"
	start, err := time.Parse(time.TimeOnly, m.StartTime)
	if err != nil {
		return false, fmt.Errorf("failed to parse start time %s: %w", m.StartTime, err)
	}

	end, err := time.Parse(time.TimeOnly, m.EndTime)
	if err != nil {
		return false, fmt.Errorf("failed to parse end time %s: %w", m.EndTime, err)
	}

	// Create a time object for comparison that only contains the time components
	// The date part is arbitrary (e.g., year=0, month=1, day=1), but must be consistent.
	tOnly := time.Date(0, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
	start = time.Date(0, 1, 1, start.Hour(), start.Minute(), start.Second(), 0, time.UTC)
	end = time.Date(0, 1, 1, end.Hour(), end.Minute(), end.Second(), 0, time.UTC)

	// Check if the current time (tOnly) falls within the start and end window
	isCurrent := !tOnly.Before(start) && tOnly.Before(end)

	return isCurrent, nil
}
