package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	el := strings.Split(input, "")
	var el2 []string
	for i := 0; i < len(el); i++ {
		if el[i] == " " {
			continue
		}
		if el[i] == "+" {
			el[i] = "/"
			el2 = append(el2, el[i])
			continue
		}
		if el[i] == "-" {
			el[i] = "/-"
		}
		el2 = append(el2, el[i])
	}

	for i := 0; i < len(el); i++ {
		if el[i] == "" {
			copy(el[i:], el[i+1:])
			el[len(el)-1] = ""
			el = el[:len(el)-1]
		}
		el2 = append(el2, el[i])
	}

	if len(el2) > 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	a, err := strconv.Atoi(el2[0])
	if err != nil {
		return "", fmt.Errorf("bad %w", err)
	}

	b, err := strconv.Atoi(el2[1])
	if err != nil {
		return "", fmt.Errorf("bad %w", err)
	}
	output = fmt.Sprint(strconv.Itoa(a + b))
	return output, nil
}
