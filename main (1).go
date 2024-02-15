package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ROMANS = map[string]int{
	"M": 1000, "D": 500, "C": 100,
	"L": 50, "X": 10, "V": 5, "I": 1}

// sum returns the sum of two integers.
//
// Parameters:
//
//	a - an integer
//	b - an integer
//
// Return type: an integer
func sum(a, b int) int {
	return a + b
}

// sub is a Go function that subtracts two integers.
//
// Parameters:
//
//	a - an integer
//	b - an integer
//
// Return type: an integer
func sub(a, b int) int {
	return a - b
}

// multy is a Go function that multiplies two integers.
//
// It takes two integer parameters and returns an integer.
func multy(a, b int) int {
	return a * b
}

// div is a Go function that takes two integers as parameters and returns an integer.
//
// The parameters are a and b, both of type int. The return type is int.
func div(a, b int) int {
	return a / b
}

// findArg finds the operator in the given line.
//
// line string - The input line to search for the operator.
// (string, error) - The operator found and an error if not found.
func findArg(line string) (string, error) {
	switch {
	case strings.Contains(line, "+"):
		return "+", nil
	case strings.Contains(line, "-"):
		return "-", nil
	case strings.Contains(line, "*"):
		return "*", nil
	case strings.Contains(line, "/"):
		return "/", nil
	default:
		return "", fmt.Errorf("can't find operator")
	}
}

// calculation performs the specified operation on two integers.
//
// It takes three parameters: a, an integer; b, an integer; op, a string.
// It returns num, an integer, and err, an error.
func calculation(a, b int, op string) (num int, err error) {
	switch op {
	case "+":
		num = sum(a, b)
	case "-":
		num = sub(a, b)
	case "*":
		num = multy(a, b)
	case "/":
		num = div(a, b)
	default:
		err = fmt.Errorf("%s not found", op)
	}

	return
}

// isRoman checks if the input string is a Roman numeral.
//
// It takes a string parameter 'num' and returns a boolean value.
func isRoman(num string) bool {
	if _, err := ROMANS[strings.Split(num, "")[0]]; err {
		return true
	}

	return false
}

// romanToInt converts a Roman numeral to an integer.
//
// It takes a string representing a Roman numeral and returns an integer.
func romanToInt(num string) int {
	sum := 0
	n := len(num)

	for i := 0; i < n; i++ {
		if i != n-1 && ROMANS[string(num[i])] < ROMANS[string(num[i+1])] {
			sum += ROMANS[string(num[i+1])] - ROMANS[string(num[i])]
			i++
			continue
		}

		sum += ROMANS[string(num[i])]
	}

	return sum
}

// intToRoman converts an integer to a Roman numeral.
//
// Parameter:
//
//	num - an integer to be converted to a Roman numeral
//
// Return type:
//
//	string - the Roman numeral representation of the input integer
func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

// getNumsAndType takes a string and an operator, splits the string by the operator,
// and returns two integers and a boolean indicating if the numbers are in Roman
// numeral format or not, along with an error if any.
func getNumsAndType(line string, op string) (a, b int, rom bool, err error) {
	nums := strings.Split(line, op)

	if len(nums) > 2 {
		return a, b, rom, fmt.Errorf("many operators")
	}

	firstRomType := isRoman(nums[0])
	secondRomType := isRoman(nums[1])

	if firstRomType != secondRomType {
		return a, b, rom, fmt.Errorf("different format")
	}

	if firstRomType && secondRomType {
		rom = true
		a = romanToInt(nums[0])
		b = romanToInt(nums[1])
	} else {
		a, err = strconv.Atoi(nums[0])
		if err != nil {
			return
		}

		b, err = strconv.Atoi(nums[1])
		if err != nil {
			return
		}
	}

	if a < 1 || a > 10 || b < 0 || b > 10 {
		return a, b, rom, fmt.Errorf("%d or %d less 0 or more 10", a, b)
	}

	return a, b, rom, nil
}

// main is the entry point of the program.
//
// It reads user input from the console and performs calculations based on the input.
// No parameters.
// No return values.
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Для выхода введите !exit\nВведите пример: ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, " ", "")

		if line == "!exit" {
			fmt.Println("exiting..")
			return
		}

		operator, err := findArg(line)
		if err != nil {
			panic(err)
		}

		a, b, isRom, err := getNumsAndType(line, operator)
		if err != nil {
			panic(err)
		}

		result, err := calculation(a, b, operator)
		if err != nil {
			panic(err)
		}

		if isRom {
			if result <= 0 {
				panic("roman numbers can't less 0")
			}

			first := intToRoman(a)
			second := intToRoman(b)
			res := intToRoman(result)

			fmt.Println(first, operator, second, "=", res)
		} else {
			fmt.Println(a, operator, b, "=", result)
		}
	}
}
