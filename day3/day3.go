package day3

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type Matcher struct {
	MatchString []byte
	MatchIndex  int
}

func (m *Matcher) CheckMatch(b byte) bool {
	if b == m.MatchString[m.MatchIndex] {
		m.MatchIndex++
		if m.MatchIndex == len(m.MatchString) {
			m.MatchIndex = 0
			return true
		}
	} else {
		m.MatchIndex = 0
	}
	return false
}

type SearchPhase string

const (
	START          SearchPhase = "START"
	FIRST_OPERAND  SearchPhase = "FIRST_OPERAND"
	SECOND_OPERAND SearchPhase = "SECOND_OPERAND"
)

func ExecuteCalculations(inputPath string, useToggle bool) (int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	doMatcher := &Matcher{MatchString: []byte("do()"), MatchIndex: 0}
	dontMatcher := &Matcher{MatchString: []byte("don't()"), MatchIndex: 0}
	mulMatcher := &Matcher{MatchString: []byte("mul("), MatchIndex: 0}
	shouldMultiply := true
	reader := bufio.NewReader(file)
	currentPhase := START
	var firstOperand = []byte{}
	var secondOperand = []byte{}
	sum := 0
	activeToggleMatcher := dontMatcher

	for {
		byte, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		if useToggle {
			hasMatch := activeToggleMatcher.CheckMatch(byte)
			if hasMatch {
				if activeToggleMatcher == dontMatcher {
					shouldMultiply = false
					activeToggleMatcher = doMatcher
				} else {
					shouldMultiply = true
					activeToggleMatcher = dontMatcher
				}
			}
		}
		if shouldMultiply {
			switch currentPhase {
			case START:
				{
					isMatched := mulMatcher.CheckMatch(byte)
					if isMatched {
						currentPhase = FIRST_OPERAND
					}
				}
			case FIRST_OPERAND:
				{
					stringValue := string(byte)
					isComma := stringValue == ","
					if isComma {
						if len(firstOperand) > 0 {
							currentPhase = SECOND_OPERAND
						} else {
							firstOperand = nil
							currentPhase = START
						}
					} else {
						_, err := strconv.Atoi(string(byte))
						if err != nil {
							firstOperand = nil
							currentPhase = START
						} else {

							firstOperand = append(firstOperand, byte)
						}
					}
				}
			case SECOND_OPERAND:
				{
					stringValue := string(byte)
					isClosingParen := stringValue == ")"
					if isClosingParen {
						if len(secondOperand) > 0 {
							currentPhase = START
							a, err := strconv.Atoi(string(firstOperand))
							if err != nil {
								return 0, err
							}
							b, err := strconv.Atoi(string(secondOperand))
							if err != nil {
								return 0, err
							}
							sum += a * b
						}
						firstOperand = nil
						secondOperand = nil
						currentPhase = START
					} else {
						_, err := strconv.Atoi(string(byte))
						if err != nil {
							firstOperand = nil
							secondOperand = nil
							currentPhase = START
						} else {
							secondOperand = append(secondOperand, byte)
						}
					}
				}
			}
		}
	}
	return sum, nil
}
