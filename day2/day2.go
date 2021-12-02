package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day2")

	Part2()
}

func Part1() {
	f, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	x := 0
	y := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		amt, _ := strconv.Atoi(line[1])
		fmt.Println(dir, " ", amt)
		if dir == "forward" {
			x += amt
		} else if dir == "down" {
			y += amt
		} else {
			y -= amt
		}
	}
	fmt.Println(x * y)
}

func Part2() {
	f, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	x := 0
	y := 0
	aim := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		amt, _ := strconv.Atoi(line[1])
		if dir == "forward" {
			x += amt
			y += aim * amt
		} else if dir == "down" {
			aim += amt
		} else {
			aim -= amt
		}
		fmt.Println("dir", dir, "amt", amt, "x", x, "y", y, "aim", aim)
	}
	fmt.Println(x * y)
}
