package math_test

import (
	"testing"

	"github.com/darmayasa221/go-basic/math"
	"github.com/stretchr/testify/assert"
)

func TestArraySlice(t *testing.T) {
	t.Run("should collection of 5 number", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3, 4, 5}
		result := 15
		// Action
		ArraySlice := math.ArraySlice(numbers)
		// Assert
		assert.Equal(t, result, ArraySlice, "should return an int %q", result)
	})
	t.Run("should collection of any size", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3}
		result := 6
		// Action
		ArraySlice := math.ArraySlice(numbers)
		// Assert
		assert.Equal(t, result, ArraySlice, "should return an int %q", result)
	})

}
