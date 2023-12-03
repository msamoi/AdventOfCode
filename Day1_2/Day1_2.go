package AdvOfCode

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Solve() int {
	file, err := os.Open("Day1_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		tmp, err := getFirstLastNumbers(text)
		if err != nil {
			log.Fatal(err)
		}
		sum += tmp
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func getFirstLastNumbers(input string) (int, error) {
	var firstNum rune
	var lastNum rune
	indexes, digits, success := findFirstLastDigits(input)
	firstWord, lastWord := getFirstLastWordIndexesAndValues(input)
	if success {
		firstNum = digits[0]
		lastNum = digits[1]
	} else {
		firstNum = rune(firstWord[1] + 48)
		lastNum = rune(lastWord[1] + 48)
		return strconv.Atoi(string([]rune{firstNum, lastNum}))
	}

	if firstWord[0] < indexes[0] {
		firstNum = rune(firstWord[1] + 48)
	}
	if lastWord[0] > indexes[1] {
		lastNum = rune(lastWord[1] + 48)
	}
	return strconv.Atoi(string([]rune{firstNum, lastNum}))
}

func getFirstLastWordIndexesAndValues(input string) ([]int, []int) {
	left, right := []int{len(input), -1}, []int{-1, -1}
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, number := range numbers {
		tmp, err := IndexOfFirstSubstring(input, number)
		if err != nil {
			continue
		}
		if tmp < left[0] {
			left[0] = tmp
			left[1] = i
		}
	}
	for i, number := range numbers {
		tmp, err := IndexOfLastSubstring(input, number)

		if err != nil {
			continue
		}
		if tmp > right[0] {
			right[0] = tmp
			right[1] = i
		}
	}
	return left, right
}

func IndexOfFirstSubstring(str, subStr string) (int, error) {
	for i := 0; i < len(str)-len(subStr); i++ {
		if str[i:i+len(subStr)] == subStr {
			return i, nil
		}
	}
	return -1, errors.New("no substring found")
}

func IndexOfLastSubstring(str, subStr string) (int, error) {
	for i := len(str) - len(subStr); i > 0; i-- {
		if str[i:i+len(subStr)] == subStr {
			return i, nil
		}
	}
	return -1, errors.New("no substring found")
}

func findFirstLastDigits(input string) ([]int, []rune, bool) {
	left := -1
	right := len(input) - 1
	runes := []rune(input)
	for i := 0; i < len(input); i++ {
		if unicode.IsDigit(runes[i]) {
			left = i
			break
		}
	}
	if left == -1 {
		return nil, nil, false
	}
	for i := right; i >= left; i-- {
		if unicode.IsDigit(runes[i]) {
			right = i
			break
		}
	}
	return []int{left, right}, []rune{runes[left], runes[right]}, true
}
