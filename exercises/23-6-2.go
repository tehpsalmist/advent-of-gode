package exercises

import (
	"fmt"
	"strconv"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day6Exercise2() {
	input := utils.LineReaderString("./inputs/23-6.txt")

	time, _ := strconv.ParseInt(strings.Join(strings.Fields(input[0])[1:], ""), 0, 64)
	distance, _ := strconv.ParseInt(strings.Join(strings.Fields(input[1])[1:], ""), 0, 64)

	total := evaluateRace2(time, distance)

	fmt.Println(total)
}

func evaluateRace2(time, distance int64) int64 {
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

	win := int64(0)
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
