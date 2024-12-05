package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCalculationsShort(t *testing.T) {
	sum, err := ExecuteCalculations("./day3_test_input.txt", false)
	assert.Nil(t, err)
	expected := 161
	assert.Equal(t, expected, sum)
}

func TestExecuteCalculationsActual(t *testing.T) {
	sum, err := ExecuteCalculations("./day3_actual_input.txt", false)
	assert.Nil(t, err)
	t.Log(sum)
}

func TestExecuteCalculationsShortWithToggle(t *testing.T) {
	sum, err := ExecuteCalculations("./day3_test_input2.txt", true)
	assert.Nil(t, err)
	expected := 48
	assert.Equal(t, expected, sum)
}

func TestExecuteCalculationsActualWithToggle(t *testing.T) {
	sum, err := ExecuteCalculations("./day3_actual_input.txt", true)
	assert.Nil(t, err)
	t.Log(sum)
}
