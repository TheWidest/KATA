package main

import (
	"bytes"
	"errors"
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	words := strings.Split(text, " ")

	if len(words) != 3 {
		panic("Неверный формат ввода")
	}
	if words[1] != "+" && words[1] != "-" && words[1] != "*" && words[1] != "/" {
		panic("Неверный формат ввода")
	}

	words[2] = strings.TrimSpace(words[2])

	if _, err := strconv.Atoi(words[0]); err == nil {
		if _, err := strconv.Atoi(words[2]); err == nil {
			Arab1, _ := strconv.Atoi(words[0])
			Arab2, _ := strconv.Atoi(words[2])			

			if Arab1 > 0 && Arab1 < 11 && Arab2 > 0 && Arab2 < 11 {
				switch words[1] {
				case "+":
					fmt.Println(Arab1 + Arab2)
				case "-":
					fmt.Println(Arab1 - Arab2)
				case "*":
					fmt.Println(Arab1 * Arab2)
				case "/":
					fmt.Println(Arab1 / Arab2)
				}
			} else {
				panic("Неверный формат ввода")
			}
		} else {
			panic("Неверный формат ввода")
		}
	} else if _, ok := RomanToInt([]byte(words[0])); ok == true {
		if _, ok := RomanToInt([]byte(words[2])); ok == true {
			Rom1, _ := RomanToInt([]byte(words[0]))
			Rom2, _ := RomanToInt([]byte(words[2]))

			if Rom1 > 0 && Rom1 < 11 && Rom2 > 0 && Rom2 < 11 {
				answer := 0
				switch words[1] {
				case "+":
					answer = Rom1 + Rom2
				case "-":
					answer = Rom1 - Rom2
				case "*":
					answer = Rom1 * Rom2
				case "/":
					answer = Rom1 / Rom2
				}
				fmt.Println(intToRoman(answer))
			} else {
				panic("Неверный формат ввода")
			}
		} else {
			panic("Неверный формат ввода")
		}
	} else {
		panic("Неверный формат ввода")
	}

}


const (
	_maxRoman = 3999
	_minRoman = 1
)

var (
	InvalidRomanNumeral = errors.New("invalid roman numeral")
	IntegerOutOfBounds  = errors.New("integer must be between 1 and 3999")
)

// numeral describes the value and symbol of a single roman numeral
type numeral struct {
	val int
	sym []byte
}

// _numerals are all unique numerals ordered from largest to smallest
var _numerals = []numeral{
	{1000, []byte("M")},
	{900, []byte("CM")},
	{500, []byte("D")},
	{400, []byte("CD")},
	{100, []byte("C")},
	{90, []byte("XC")},
	{50, []byte("L")},
	{40, []byte("XL")},
	{10, []byte("X")},
	{9, []byte("IX")},
	{5, []byte("V")},
	{4, []byte("IV")},
	{1, []byte("I")},
}

// lookup arrays used for converting from an int to a roman numeral extremely quickly.
// method from here: https://rosettacode.org/wiki/Roman_numerals/Encode#Go
var (
	m0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	m1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	m2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	m3 = []string{"", "M", "MM", "MMM"}
)

// IntToString converts an integer value to a roman numeral string. An error is
// returned if the integer is not between 1 and 3999.
func IntToString(input int) (string, error) {
	if outOfBounds(input) {
		return "", IntegerOutOfBounds
	}
	return intToRoman(input), nil
}

// IntToBytes converts an integer value to a roman numeral byte array. An error is
// returned if the integer is not between 1 and 3999.
func IntToBytes(input int) ([]byte, error) {
	str, err := IntToString(input)
	return []byte(str), err
}

// outOfBounds checks to ensure an input value is valid for roman numerals without the need of
// vinculum (used for values of 4,000 and greater)
func outOfBounds(input int) bool {
	return input < _minRoman || input > _maxRoman
}

func intToRoman(n int) string {
	// This is efficient in Go. The 4 operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return m3[n%1e4/1e3] + m2[n%1e3/1e2] + m1[n%100/10] + m0[n%10]
}

// StringToInt converts a roman numeral string to an integer. Roman numerals for numbers
// outside of the range 1 to 3,999 will return an error. Empty strings will return 0
// with no error thrown.
func StringToInt(input string) (int, error) {
	return BytesToInt([]byte(input))
}

// BytesToInt converts a roman numeral byte array to an integer. Roman numerals for numbers
// outside of the range 1 to 3,999 will return an error. Nil or empty []byte will return 0
// with no error thrown.
func BytesToInt(input []byte) (int, error) {
	if input == nil || len(input) == 0 {
		return 0, nil
	}
	if output, ok := RomanToInt(input); ok {
		return output, nil
	}
	return 0, InvalidRomanNumeral
}

func RomanToInt(input []byte) (int, bool) {
	var output int
	for _, rom := range _numerals {
		for bytes.HasPrefix(input, rom.sym) {
			output += rom.val
			input = input[len(rom.sym):]
		}
	}
	// If we are still left with input string values then the
	// input was invalid and the bool is returned as False
	return output, len(input) == 0
}