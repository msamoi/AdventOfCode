package AdvOfCode

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Solve() int {
	file, err := os.Open("Day1/input.txt")
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
		tmp, err := getFirstLastDigits(text)
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

func getFirstLastDigits(input string) (int, error) {
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
		return 0, nil
	}
	for i := right; i >= left; i-- {
		if unicode.IsDigit(runes[i]) {
			right = i
			break
		}
	}
	out := string([]rune{runes[left], runes[right]})
	return strconv.Atoi(out)
}
