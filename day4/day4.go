package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	val int
	on  bool
}

type Card [][]*Square

var moves []int
var cards []Card

func main() {
	fmt.Println("Day4")

	f, err := os.Open("day4.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, e := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(e)
		moves = append(moves, i)
	}

	fmt.Println(moves)

	createCards(lines[2:])

	fmt.Println(cards)

	part1()
}

func createCards(lines []string) {
	//fmt.Println(lines)
	var card = make(Card, 5)
	ci := 0
	for _, e := range lines {
		if e == "" {
			ci = 0
			cards = append(cards, card)
			card = make(Card, 5)
			//fmt.Println("new card at ", i)
			continue
		}
		//fmt.Println("A new row at ", ci, card[ci])
		for _, s := range strings.Fields(e) {
			v, _ := strconv.Atoi(s)
			card[ci] = append(card[ci], &Square{v, false})
		}
		//fmt.Println("A finished row at ", ci, card[ci])
		ci++
	}
	cards = append(cards, card)
}

func part1() {
	var winner Card
	for _, m := range moves {
		if winner != nil {
			break
		}
		for _, c := range cards {
			makeMove(&c, m)
			if checkIfWinner(c) {
				winner = c
				score := m * sumUnused(c)
				fmt.Println("Score", score)
				break
			}
		}

	}
	fmt.Println("Winner", winner)
}

func makeMove(card *Card, move int) {
	for i, row := range *card {
		for j, val := range row {
			if val.val == move {
				val.on = true
				fmt.Println("Found", move, "at", i, j, card)
			}
		}
	}
}

func checkIfWinner(card Card) bool {
	completeRow := true
	for _, row := range card {
		for _, val := range row {
			//fmt.Println("Checking row", row, val)
			if !val.on {
				completeRow = false
				break
			}
		}
		if completeRow {
			fmt.Println("Found winning row", row)
			printCard(card)
			return true
		}
		completeRow = true
	}
	completeCol := true
	for j := 0; j < 5; j++ {
		for _, row := range card {
			if !row[j].on {
				completeCol = false
				break
			}
		}
		if completeCol {
			fmt.Println("Found winning col", j)
			printCard(card)
			return true
		}
		completeCol = true
	}
	return false
}

func sumUnused(card Card) int {
	sum := 0
	for _, row := range card {
		for _, val := range row {
			if !val.on {
				sum += val.val
			}
		}
	}
	return sum
}

func printCard(c Card) {
	fmt.Println("Print Card")
	for _, row := range c {
		for _, val := range row {
			fmt.Print(*val, " ")
		}
		fmt.Println()
	}
}
