package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareLocationListsShortInput(t *testing.T) {
	expected := 11
	actual, err := CompareLocationLists("./day1_test_input.txt")
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, expected, actual, "Expected the sum of the absolute difference between the two lists to be 12")
}

func TestCompareLocationListsActual(t *testing.T) {
	diff, err := CompareLocationLists("./day1_input.txt")
	assert.Nil(t, err, "Expected no error")
	fmt.Println(fmt.Sprintf("Difference between lists was: %v", diff))
	assert.NotNil(t, diff)
}

func TestCalculateSimilarityShortInput(t *testing.T) {
	expected := 31
	actual, err := CalculateSimilarity("./day1_test_input.txt")
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, expected, actual, "Expected the sum of the absolute difference between the two lists to be 12")
}

func TestCalculateSimilarityActual(t *testing.T) {
	similarity, err := CalculateSimilarity("./day1_input.txt")
	assert.Nil(t, err, "Expected no error")
	fmt.Println(fmt.Sprintf("Similarity between lists was: %v", similarity))
}
