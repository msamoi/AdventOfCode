package AdvOfCode

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type number struct {
	y      int
	x      int
	length int
}

var numberMap = make(map[string]number)

func Solve() int {
	file, err := os.Open("Day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return sumPartNumbers(lines)
}

func sumPartNumbers(input []string) int {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			curr := input[y][x]
			if (curr >= '0' && curr <= '9') || curr == '.' {
				continue
			}

			findSurroundingNumbers(input, y, x)
		}
	}

	out := 0
	for _, entry := range numberMap {
		tmp, err := strconv.Atoi(input[entry.y][entry.x : entry.x+entry.length])
		if err != nil {
			log.Fatal(err)
		}
		out += tmp
	}
	return out
}

func findSurroundingNumbers(input []string, yIndex int, xIndex int) {
	y := []int{-1, 1}
	x := []int{-1, 1}
	switch yIndex {
	case 0:
		y[0] = 0
	case len(input):
		y[1] = 0
	}

	switch xIndex {
	case 0:
		x[0] = 0
	case len(input[yIndex]) - 1:
		x[1] = 0
	}

	numeric := regexp.MustCompile(`\d+`)
	for i := y[0]; i <= y[1]; i++ {
		begin := xIndex + x[0]
		end := xIndex + x[1]
		toCheck := input[yIndex+i][begin : end+1]
		numIndex := numeric.FindAllStringIndex(toCheck, -1)
		if numIndex == nil {
			continue
		}

		for _, index := range numIndex {
			tmp := getNumberCoords(input[yIndex+i], begin+index[0])
			tmp.y = yIndex + i

			key := strconv.Itoa(tmp.y) + "-" + strconv.Itoa(tmp.x)

			numberMap[key] = tmp
		}

	}
}

func getNumberCoords(input string, index int) number {

	var out number
	i := 0
	for i = index; i >= 0; i-- {
		if input[i] >= '0' && input[i] <= '9' {
			continue
		}
		break
	}
	out.x = i + 1

	for i = out.x; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			continue
		}
		break
	}
	out.length = i - out.x

	return out
}
