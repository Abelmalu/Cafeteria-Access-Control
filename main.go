package main

import (
	"fmt"
	"strings"
	"time"
)

// The target format we expect from the user, e.g., "12:30 PM" or "8:00 AM"
// NOTE: The space between "04" (minutes) and "PM" is CRUCIAL for matching inputs like "1:45 PM".
const userTimeLayout = "3:04 PM"

// The time format MySQL TIME expects (HH:MM:SS)
const mysqlTimeFormat = "15:04:05"

// ParseUserTime converts a user-provided time string (e.g., "1:30 PM")
// into a standard time.Time object.
func ParseUserTime(timeStr string) (time.Time, error) {
	// 1. Clean the input (trim spaces). We rely on the layout to handle AM/PM case.
	cleanedStr := strings.TrimSpace(timeStr)

	// 2. Parse the string using the defined layout.
	t, err := time.Parse(userTimeLayout, cleanedStr)
	if err != nil {
		// Return a more descriptive error if parsing fails
		return time.Time{}, fmt.Errorf("could not parse time '%s'. Expected format like '1:30 PM': %w", timeStr, err)
	}

	// The resulting `t` object holds the correct 24-hour time.
	return t, nil
}

// FormatTimeForMySQL formats the time.Time object into the HH:MM:SS string
// required by MySQL's TIME type.
func FormatTimeForMySQL(t time.Time) string {
	return t.Format(mysqlTimeFormat)
}

func main() {
	// --- User Input ---
	userInput := "1:45 PM"
	userInputNoon := "12:00 PM"
	userInputMorning := "2:15 PM"

	// Test Case 1: 1:45 PM
	parsedTime1, err1 := ParseUserTime(userInput)
	if err1 != nil {
		fmt.Println("Error 1:", err1)
		return
	}
	fmt.Printf("Input: %s\n", userInput)
	fmt.Printf("   -> MySQL Format: %s (Go 24h: %s)\n\n", FormatTimeForMySQL(parsedTime1), parsedTime1.Format("15:04"))

	// Test Case 2: 12:00 PM (Noon)
	parsedTime2, err2 := ParseUserTime(userInputNoon)
	if err2 != nil {
		fmt.Println("Error 2:", err2)
		return
	}
	fmt.Printf("Input: %s\n", userInputNoon)
	fmt.Printf("   -> MySQL Format: %s (Go 24h: %s)\n\n", FormatTimeForMySQL(parsedTime2), parsedTime2.Format("15:04"))

	// Test Case 3: 8:15 AM
	parsedTime3, err3 := ParseUserTime(userInputMorning)
	if err3 != nil {
		fmt.Println("Error 3:", err3)
		return
	}
	fmt.Printf("Input: %s\n", userInputMorning)
	fmt.Printf("   -> MySQL Format: %s (Go 24h: %s)\n", FormatTimeForMySQL(parsedTime3), parsedTime3.Format("15:04"))
}
