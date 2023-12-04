package exercises

import (
	"fmt"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day4Exercise2() {
	total := 0

	// Assuming that all candidate lists are unique, maximum matches will be 10.
	// Size 11 slice allows us to deque the first and still have room in the slice
	// to avoid index out of range problems. Values initialize at 1 because we always
	// start with one copy of that scratchcard to evaluate.
	copyQueue := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	lineCount := 0
	utils.LineReaderStringIterator("./23-4.txt", func(line string) {
		lineCount++
		// split line by |
		parts := strings.Split(line, "|")

		// split first part for winners
		winners := map[string]bool{}

		for _, w := range strings.Fields(strings.Split(parts[0], ":")[1]) {
			winners[w] = true
		}

		// split second part for candidates
		candidates := strings.Fields(parts[1])

		// simple dequeing
		copies := copyQueue[0]
		copyQueue = append(copyQueue[1:], 1)

		winningMatches := 0
		for _, c := range candidates {
			if _, ok := winners[c]; ok {
				winningMatches++
			}
		}

		// the queue represents the upcoming scratchcards, so we can iterate through them for as many matches as we found
		// and increment the copies of each by as many copies we have of the current card
		for i := 0; i < winningMatches; i++ {
			copyQueue[i] += copies
		}

		total += copies
	})

	fmt.Println(total)
}
