package main

import "testing"

func TestGetMax(t *testing.T) {
	data := []int{100, 1, 10, 200, 3, 4, 900}
	result := GetMax(data)
	if result != 900 {
		t.Error("Expected: {} Got: {}", 900, result)
	}

}
