package day1

import (
	"bufio"
	"errors"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ProcessInput(inputPath string) ([]float64, []float64, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var listA []float64
	var listB []float64

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) != 2 {
			return nil, nil, errors.New("Unexpected input")
		}
		a, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			return nil, nil, err
		}
		b, err := strconv.ParseFloat(values[1], 64)
		if err != nil {
			return nil, nil, err
		}
		listA = append(listA, a)
		listB = append(listB, b)
	}

	err = scanner.Err()
	if err != nil {
		return nil, nil, err
	}

	return listA, listB, nil
}

func CompareLocationLists(inputPath string) (int, error) {
	var diff float64
	listA, listB, err := ProcessInput(inputPath)
	if err != nil {
		return 0, err
	}
	sort.Float64s(listA)
	sort.Float64s(listB)

	for i := 0; i < len(listA); i++ {
		diff += math.Abs(listA[i] - listB[i])
	}

	return int(diff), nil
}

func CalculateSimilarity(inputPath string) (int, error) {
	var similarity float64
	listA, listB, err := ProcessInput(inputPath)
	if err != nil {
		return 0, err
	}
	freqMap := map[float64]int{}
	for _, v := range listB {
		freqMap[v]++
	}
	for _, v := range listA {
		similarity += v * float64(freqMap[v])
	}
	return int(similarity), nil
}
