package day2

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func CountUnsafeReports(inputPath string, isSafe func(values []int, replacesAllowed int) bool) (int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		intValues := make([]int, len(values))
		for i, v := range values {
			intValues[i], err = strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
		}
		safe := isSafe(intValues, 0)
		if safe {
			safeCount++
		}
	}

	err = scanner.Err()
	if err != nil {
		return 0, err
	}

	return safeCount, nil
}

func isSafeV1(values []int, _ int) bool {
	if len(values) == 0 {
		return false
	}
	if len(values) == 1 {
		return true
	}
	prevValue := values[0]
	prevDirection := 0
	for i := 1; i < len(values); i++ {
		currValue := values[i]
		diff := currValue - prevValue
		if diff == 0 {
			return false
		}
		currentDirection := 0
		if diff < 0 {
			currentDirection = -1
		} else {
			currentDirection = 1
		}
		if prevDirection == 0 {
			prevDirection = currentDirection
		} else if prevDirection != currentDirection {
			return false
		}
		absDiff := math.Abs(float64(diff))
		if absDiff < 1 || absDiff > 3 {
			return false
		}
		prevValue = currValue
	}
	return true
}

func isSafeV2(values []int, replacesAllowed int) bool {
	safe := isSafeV1(values, 0)
	if safe {
		return true
	}
	for i := 0; i < len(values); i++ {
		testSet := []int{}
		for j, v := range values {
			if j != i {
				testSet = append(testSet, v)
			}
		}
		safe = isSafeV1(testSet, 0)
		if safe {
			return true
		}
	}
	return safe
}
