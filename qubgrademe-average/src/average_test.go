package main

import "testing"

// Ensure that function returns desired output
func TestAverageMarks_1(t *testing.T) {
	marks := map[string]int{"m1": 1, "m2": 5, "m3": 10, "m4": 20, "m5": 5}
	avg, err := averageMarks(marks)

	if err != nil {
		t.Error("Error: expected no error, got the following:", err.Error())
	}

	if avg != 8.2 {
		t.Errorf("Average: expected '8.2', got '%f'", avg)
	}
}

// Ensure that functions returns desired output
func TestAverageMarks_2(t *testing.T) {
	marks := map[string]int{"m1": 1, "m2": 99, "m3": 99, "m4": 99, "m5": 99}
	avg, err := averageMarks(marks)

	if err != nil {
		t.Error("Error: expected no error, got the following:", err.Error())
	}

	if avg != 79.4 {
		t.Errorf("Average: expected '79.4', got '%f'", avg)
	}
}

// Ensure that function returns 0 when all mark values are 0
func TestAverageMarks_3(t *testing.T) {
	marks := map[string]int{"m1": 0, "m2": 0, "m3": 0, "m4": 0, "m5": 0}
	avg, err := averageMarks(marks)

	if err != nil {
		t.Error("Error: expected no error, got the following:", err.Error())
	}

	if avg != 0 {
		t.Errorf("Average: expected '0', got '%f'", avg)
	}
}

// Ensure that function returns error when marks map is nil
func TestAverageMarks_4(t *testing.T) {
	avg, err := averageMarks(nil)

	if err == nil {
		t.Error("Error: expected an error, got the following:", err.Error())
	}

	if avg != -1 {
		t.Errorf("Average: expected '-1', got '%f'", avg)
	}
}

// Ensure that function returns error when marks map is empty
func TestAverageMarks_5(t *testing.T) {
	marks := map[string]int{}
	avg, err := averageMarks(marks)

	if err == nil {
		t.Error("Error: expected an error, got the following:", err.Error())
	}

	if avg != -1 {
		t.Errorf("Average: expected '-1', got '%f'", avg)
	}
}

// Ensure that function returns error when marks map contains -13
func TestAverageMarks_6(t *testing.T) {
	marks := map[string]int{"m1": 1, "m2": -13, "m3": 10, "m4": 20, "m5": 5}
	avg, err := averageMarks(marks)

	if err == nil {
		t.Error("Error: expected an error, got the following:", err.Error())
	}

	if avg != -1 {
		t.Errorf("Average: expected '-1', got '%f'", avg)
	}
}

// Ensure that function returns error when marks map contains 998
func TestAverageMarks_7(t *testing.T) {
	marks := map[string]int{"m1": 1, "m2": 998, "m3": 10, "m4": 20, "m5": 5}
	avg, err := averageMarks(marks)

	if err == nil {
		t.Error("Error: expected an error, got the following:", err.Error())
	}

	if avg != -1 {
		t.Errorf("Average: expected '-1', got '%f'", avg)
	}
}
