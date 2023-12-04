package AdvOfCode

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	file, err := os.Open("Day2/input.txt")
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
		tmp, possible := checkGamePossibility(scanner.Text())
		if possible {
			sum += tmp
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func checkGamePossibility(input string) (int, bool) {
	stringSlice := strings.Split(input, " ")
	gameId, err := strconv.Atoi(strings.Trim(stringSlice[1], ":"))
	if err != nil {
		log.Fatal(err)
	}
	for i, word := range stringSlice {
		word = strings.Trim(word, ";,")
		switch word {
		case "red":
			if tooManyCubes(stringSlice[i-1], 12) {
				return gameId, false
			}
		case "green":
			if tooManyCubes(stringSlice[i-1], 13) {
				return gameId, false
			}
		case "blue":
			if tooManyCubes(stringSlice[i-1], 14) {
				return gameId, false
			}
		}
	}
	return gameId, true
}

func tooManyCubes(input string, amount int) bool {
	tmp, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if tmp > amount {
		return true
	}
	return false
}
