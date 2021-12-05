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

type Card struct {
	squares [][]*Square
	won     bool
}

var moves []int
var cards []*Card

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

	//fmt.Println(cards)
	for i, c := range cards {
		fmt.Println("Card", i, &c)
		printCard(*c)
	}

	//part1()
	part2()
}

func createCards(lines []string) {
	//fmt.Println(lines)
	var card = &Card{}
	card.squares = make([][]*Square, 5)
	card.won = false
	fmt.Println("First card at ", &card)
	ci := 0
	for i, e := range lines {
		if e == "" {
			ci = 0
			cards = append(cards, card)
			card = &Card{}
			card.squares = make([][]*Square, 5)
			card.won = false
			fmt.Println("new card at ", i, card)
			continue
		}
		fmt.Println("A new row at ", ci, card.squares[ci])
		for _, s := range strings.Fields(e) {
			v, _ := strconv.Atoi(s)
			card.squares[ci] = append(card.squares[ci], &Square{v, false})
		}
		//fmt.Println("A finished row at ", ci, card[ci])
		ci++
	}
	cards = append(cards, card)
}

func part1() {
	var winner *Card
	for _, m := range moves {
		if winner.won {
			break
		}
		for _, c := range cards {
			makeMove(c, m)
			if checkIfWinner(*c) {
				winner = c
				winner.won = true
				score := m * sumUnused(*c)
				fmt.Println("Score", score)
				break
			}
		}

	}
	fmt.Println("Winner", winner)
}

func makeMove(card *Card, move int) {
	for i, row := range card.squares {
		for j, val := range row {
			if val.val == move {
				val.on = true
				fmt.Println("Found", move, "at", i, j)
			}
		}
	}
}

func checkIfWinner(card Card) bool {
	completeRow := true
	for _, row := range card.squares {
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
		for _, row := range card.squares {
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
	for _, row := range card.squares {
		for _, val := range row {
			if !val.on {
				sum += val.val
			}
		}
	}
	return sum
}

func printCard(c Card) {
	fmt.Println("Print Card, won:", c.won)
	for _, row := range c.squares {
		for _, val := range row {
			fmt.Print(*val, " ")
		}
		fmt.Println()
	}
}

func part2() {
	var lastWinner *Card
	for _, m := range moves {
		if lastWinner != nil {
			break
		}
		for i, c := range cards {
			if c.won {
				fmt.Println(i, "has already won")
				continue
			}
			makeMove(c, m)
			if checkIfWinner(*c) {
				c.won = true
				fmt.Println("A card has won ", c.won)
				if checkIfLastWinner() {
					lastWinner = c
					score := m * sumUnused(*c)
					fmt.Println("Score", score)
					break
				}
			}
		}

	}
	fmt.Println("Last Winner")
	if lastWinner != nil {
		printCard(*lastWinner)
	}
}

func checkIfLastWinner() bool {
	count := 0
	for i, c := range cards {
		fmt.Println("checkLastWinner", i, c.won, c.squares[0][0])
		if c.won {
			count++
		}
	}
	fmt.Println("Number of winning cards", count)
	return count == len(cards)
}
