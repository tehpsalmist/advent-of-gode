package exercises

import (
	"fmt"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day6Exercise1() {
	total := 1

	input := utils.LineReaderString("./inputs/23-6.txt")

	times := utils.MapStringToInt64(strings.Fields(input[0])[1:])
	distances := utils.MapStringToInt64(strings.Fields(input[1])[1:])

	for i, time := range times {
		distance := distances[i]

		total *= evaluateRace(time, distance)
	}

	fmt.Println(total)
}

func evaluateRace(time, distance int64) int {
	// the best time will be the number(s) at the midpoint of the allotted time
	hold := time / 2
	running := hold
	isOdd := time%2 == 1

	if isOdd {
		// divide by 2 drops the remainder, add it back to one of them
		running++
	}

	// how far did we go?
	product := hold * running

	win := 0
	for product > distance {
		// log the win
		win++

		// see if holding it 1ms less will also win
		hold--
		running++
		product = hold * running
	}

	// the inverse of the hold/run values are also wins, so double it
	win *= 2

	if !isOdd {
		// even numbers will have a number multiplied by itself in the list, which has no inverse, so drop it
		win--
	}

	return win
}
