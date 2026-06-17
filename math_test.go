package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Soma(3, 5)
	expected := 8

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	} else {
		fmt.Printf("Test passed: Sum(3, 5) = %d\n", result)
	}
}
