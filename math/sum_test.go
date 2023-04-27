package math_test

import (
	"testing"

	"github.com/darmayasa221/go-basic/math"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	// Arrange
	numbers := [5]int{1, 2, 4, 6, 5}
	result := 18
	// Action
	sum := math.Sum(numbers)
	// Assert
	assert.Equal(t, result, sum, "should be sum all number on array")
}
