package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day1")

	//Part1()
	Part2()
}

func Part1() {
	f, err := os.Open("day1.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	prev := -1
	count := 0
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())

		if prev > 0 && current > prev {
			count++
		}
		str := fmt.Sprintln(current, " ", prev)
		fmt.Println(str)
		prev = current
	}
	fmt.Println(count)
}

func Part2() {
	f, err := os.Open("day1.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	first := -1
	second := -1
	prevSum := -1
	count := 0
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		str := fmt.Sprintln(first, " ", second, " ", current)
		fmt.Println(str)
		if second == -1 {
			second = current
			continue
		}
		if first == -1 {
			first = second
			second = current
			continue
		}
		currentSum := first + second + current
		fmt.Println(prevSum)
		if prevSum > 0 && currentSum > prevSum {
			count++
		}
		prevSum = currentSum
		first = second
		second = current
	}
	fmt.Println(count)
}
