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
func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return "", errorEmptyInput
	}
	if len(strings.TrimSpace(input)) < 3 {
		return "", errorNotTwoOperands
	}

	result := 0.0
	for _, item := range input {
		c := strings.Split(string(item), " ")
		num1, num2, err := parseArgs(c)
		if err != nil {
			return "", err
		}

		switch c[1] {
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		}
	}
	return fmt.Sprintf("%f", result), nil
}
