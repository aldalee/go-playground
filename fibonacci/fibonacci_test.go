package main

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	ch := fibonacci(10)

	for _, want := range expected {
		if got := <-ch; got != want {
			assert.Equal(t, want, got)
		}
	}
}
