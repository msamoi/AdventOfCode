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
		sum += getPowerOfSet(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func getPowerOfSet(input string) int {
	redMax, greenMax, blueMax := 0, 0, 0
	stringSlice := strings.Split(input, " ")
	for i, word := range stringSlice {
		word = strings.Trim(word, ";,")
		switch word {
		case "red":
			redMax = compareMax(stringSlice[i-1], redMax)
		case "green":
			greenMax = compareMax(stringSlice[i-1], greenMax)
		case "blue":
			blueMax = compareMax(stringSlice[i-1], blueMax)
		}
	}
	return redMax * greenMax * blueMax
}

func compareMax(curr string, prev int) int {
	currInt, err := strconv.Atoi(curr)
	if err != nil {
		log.Fatal(err)
	}
	if currInt > prev {
		return currInt
	}
	return prev
}
