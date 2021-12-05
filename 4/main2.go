package main

import (
    "strconv"
	"bufio"
	"fmt"
	"os"
    "strings"
)

type BingoCard struct {
    numbers [][]int
    unmarked: int
}


func main() {
	part1()
}


func part1() {
	file, _ := os.Open("4/bingo")
	defer file.Close()

	scanner := bufio.NewScanner(file)
    scanner.Scan()
    numbers := ints(strings.Split(scanner.Text(), ","))
    scanner.Scan()

    cards := []BingoCard{}

    for scanner.Scan() {
        cards = append(cards, scanCard(scanner))
    }

    // draw: for _, drawn := range numbers {
    //     for _, card := range cards {
    //         for rowId, row := range card {
    //             for columnId, number := range row {
    //                 if number == drawn {
    //                     card[rowId][columnId] = -1

    //                     if isWinner(card) {
    //                         fmt.Printf("Day 4.1: %d\n", sum(card) * drawn)

    //                         break draw
    //                     }
    //                 }
    //             }
    //         }
    //     }
    // }
}

func sum(card BingoCard) int {
    sum := 0

    for _, row := range card {
        for _, number := range row {
            if number == -1 {
                continue
            }
            
            sum += number
        }
    }

    return sum
}

func isWinner(card BingoCard) bool {
    for _, row := range card {
        for i, n := range row {
            if n != -1 {
               break
            }

            if i == len(row) - 1 {
                return true
            }
        }
    }

    for column := 0; column < len(card[0]); column++ {
        for row := 0; row < len(card); row++ {
            if card[row][column] != -1 {
                break
            }

            if row == len(card) - 1 {
                return true
            }
        }
    }

    return false
}

func scanCard(scanner *bufio.Scanner) BingoCard {
    card := BingoCard{}
    card = append(card, ints(strings.Fields(scanner.Text())))
    for scanner.Scan() {
        if scanner.Text() == "" {
            break
        }
        card = append(card, ints(strings.Fields(scanner.Text())))
    }

    return card
}

func ints(s []string) ([]int, int) {
    sum := 0
    numbers := make([]int, len(s))
    for i, value := range s {
        n, _ := strconv.Atoi(value)
        
        numbers[i] = n
        sum += n
    }

    return numbers, sum
}
