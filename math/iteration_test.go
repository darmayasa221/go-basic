package math_test

import (
	"testing"

	"github.com/darmayasa221/go-basic/math"
	"github.com/stretchr/testify/assert"
)

func TestIteration(t *testing.T) {
	// Arrange
	param := "a"
	result := "aaaaa"
	// Action
	iteration := math.Iteration(param)
	// Assert
	assert.Equal(t, result, iteration, "param should be iterat 5x ")
}
