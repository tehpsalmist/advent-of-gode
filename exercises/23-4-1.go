package exercises

import (
	"fmt"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day4Exercise1() {
	total := 0

	utils.LineReaderStringIterator("./23-4.txt", func(line string) {
		// split line by |
		parts := strings.Split(line, "|")

		// split first part for winners
		winners := map[string]bool{}

		for _, w := range strings.Fields(strings.Split(parts[0], ":")[1]) {
			winners[w] = true
		}

		// split second part for candidates
		candidates := strings.Fields(parts[1])

		// algo time
		points := 0

		for _, c := range candidates {
			if _, ok := winners[c]; ok {
				if points == 0 {
					points++
				} else {
					points = points * 2
				}
			}
		}

		total += points
	})

	fmt.Println(total)
}
