package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountUnsafeReports(t *testing.T) {
	count, err := CountUnsafeReports("./day2_test_input.txt", isSafeV1)
	assert.Nil(t, err)
	expected := 2
	assert.Equal(t, expected, count)
}

func TestCountUnsafeReportsActual(t *testing.T) {
	count, err := CountUnsafeReports("./day2_actual_input.txt", isSafeV1)
	assert.Nil(t, err)
	t.Logf("Safe reports: %d", count)
}

func TestIsSafeV1HandlesStart(t *testing.T) {
	testInput := []int{0, 2, 1, 0}
	safe := isSafeV2(testInput, 1)
	assert.True(t, safe)
}

func TestCountUnsafeReportsV2(t *testing.T) {
	count, err := CountUnsafeReports("./day2_test_input.txt", isSafeV2)
	assert.Nil(t, err)
	expected := 4
	assert.Equal(t, expected, count)
}

func TestCountUnsafeReportsActualV2(t *testing.T) {
	count, err := CountUnsafeReports("./day2_actual_input.txt", isSafeV2)
	assert.Nil(t, err)
	t.Logf("Safe reports: %d", count)
}
