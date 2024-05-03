package main

import (
	"errors"
	"fmt"
	"math"
)

func averageMarks(marks map[string]int) (float64, error) {
	// Ensure marks is not nil
	if marks == nil {
		return -1, errors.New("'Marks' map is nil")
	}

	// Ensure marks is not empty
	if len(marks) == 0 {
		return -1, errors.New("'Marks' map does not contain any values")
	}

	sum := 0.

	for _, value := range marks {
		// Ensure value is valid
		if value < 0 || value > 100 {
			msg := fmt.Sprintf("%d is out of range", value)
			return -1, errors.New(msg)
		}

		sum += float64(value)
	}

	if sum == 0 {
		return 0, nil
	}

	average := sum / float64(len(marks))

	average = math.Round(average*100) / 100

	return average, nil
}
