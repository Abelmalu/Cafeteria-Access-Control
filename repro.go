package main

import (
	"fmt"
	"time"
)

func main() {
	inputs := []string{
		"12:00:00",
		"12:00",
		"12:00:00 PM", // This one is still expected to fail with the current fix, but checking behavior
	}

	layouts := []string{"15:04:00", "15:04"}

	fmt.Printf("Testing robust parsing logic...\n")
	for _, input := range inputs {
		var t time.Time
		var err error
		parsed := false

		for _, layout := range layouts {
			t, err = time.Parse(layout, input)
			if err == nil {
				parsed = true
				break
			}
		}

		if parsed {
			fmt.Printf("Input '%s' PASSED: %v\n", input, t)
		} else {
			fmt.Printf("Input '%s' FAILED: All layouts failed\n", input)
		}
	}
}
