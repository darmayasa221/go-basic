package output_test

import (
	"testing"

	"github.com/darmayasa221/go-basic/output"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	// Arrange
	var param string = "hello world"
	// Action
	output := output.Output(param)
	// Assert
	assert.Equal(t, param, output, "output should be a string and have a same value")
}

func TestOutputInt(t *testing.T) {
	// Arrang
	var param int = 101
	// Action
	output := output.OutputInt(param)
	// Assert
	assert.Equal(t, param, output, "output should be a int and have same value")
}
