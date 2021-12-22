package main

import "fmt"

func main() {
	part1()
    part2()
}

func part1() {
	scoreA := 0
	positionA := 4
	scoreB := 0
	positionB := 8
	dice := 1
	rolled := 0

    isA := game1(&scoreA, &positionA, &scoreB, &positionB, true, &dice, &rolled)

    if isA {
        fmt.Println(rolled * scoreB)
    } else {
        fmt.Println(rolled * scoreA)
    }
}

func game1(scoreA *int, positionA *int, scoreB *int, positionB *int, isA bool, dice *int, rolled *int) bool {
	roll := func() int {
		value := *dice

		(*dice)++
		if *dice > 100 {
			*dice = 1
		}

		(*rolled)++

		return value
	}

	update := func(moves int, score *int, position *int) bool {
		for i := 0; i < moves; i++ {
			*position += 1

			if *position == 11 {
				*position = 1
			}
		}

		*score += *position

		if *score >= 1000 {
			return true
		}
		return false
	}

	var score *int
	var position *int

	if isA {
		score = scoreA
		position = positionA
	} else {
		score = scoreB
		position = positionB
	}

	moves := roll() + roll() + roll()

	if update(moves, score, position) {
		return isA
	}

	return game1(scoreA, positionA, scoreB, positionB, !isA, dice, rolled)
}

// Because we are grouping for performance reasons but it does mean when the
// players throws one of the numbers it wins in multiple universes with that
// throw if he wins.
// Counts are calculated by triple for loop from 1 - 3.
var counts = map[int]int{3:1, 4:3, 5:6, 6:7, 7:6, 8:3, 9:1,}
var sheet = &map[bool]int{}

func part2() {
    // Person rolls three times the dice with values 1 to 3. Meaning that
    // person can roll 3*3*3 possible combinations which is 27. But in this
    // game rolling for example: 1,2,3 is the same as 3,2,1, because both mean
    // we move 6 spots.
    //
    // So technically a person can roll: 3,4,5,6,7,8,9 total. 

    for i := 3; i <= 9; i++ {
        game2(true, 0, 4, 0, 8, i, counts[i]) 
    }

    // Print stats of both players. Too lazy to find only the highest value.
    fmt.Println(sheet)
}

func game2(isA bool, score int, position int, otherScore int, otherPosition int, move int, rolls int) {
    position += move
    if position >= 11 {
        position -= 10
    }
    score += position

    if score >= 21 {
        (*sheet)[isA] += rolls
    } else {
        for i := 3; i <= 9; i++ {
            game2(!isA, otherScore, otherPosition, score, position, i, rolls * counts[i]) 
        }
    }
}
