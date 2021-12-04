package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day3")

	Part2()
}

func Part2() {
	f, err := os.Open("sample.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	matches := make([]string, len(lines))
	copy(matches, lines)
	for i := range lines[0] {
		if len(matches) == 1 {
			break
		}
		max := getMaxNumber(matches, i)
		fmt.Println(i, max)
		for _, m := range matches {

		}
	}

	str := ""
	o2, _ := strconv.ParseInt(str, 2, 64)
	co2, _ := strconv.ParseInt(invert(str), 2, 64)

	fmt.Println(matches, o2, co2, o2*co2)
}

func Part1() {
	f, err := os.Open("day3.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	str := ""
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := range lines[0] {
		str += getMaxNumber(lines, i)
	}

	gamma, _ := strconv.ParseInt(str, 2, 64)
	epsilon, _ := strconv.ParseInt(invert(str), 2, 64)

	fmt.Println(str, gamma, epsilon, gamma*epsilon)
}

func invert(str string) string {
	inverted := ""
	for _, ch := range str {
		if string(ch) == "1" {
			inverted += "0"
		} else {
			inverted += "1"
		}
	}
	return inverted
}

func getMaxNumber(lines []string, index int) string {
	numZeros := 0
	numOnes := 0
	for _, line := range lines {

		c := string(line[index])
		//fmt.Println(c)
		if c == "0" {
			numZeros++
		} else {
			numOnes++
		}
	}
	if numOnes > numZeros {
		return "1"
	}
	return "0"
}
