package exercises

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"8thday.dev/aog/utils"
)

func Year23Day3Exercise1() {
	total := int64(0)

	schematic := utils.LineReaderString("./23-3.txt")

	for y, line := range schematic {
		currentNumber := []rune{}

		for x, char := range line {
			isDigit := unicode.IsDigit(char)

			if isDigit {
				currentNumber = append(currentNumber, char)
			}

			if (x == len(line)-1 && isDigit) || (len(currentNumber) > 0 && !isDigit) {
				numEndIndex := x

				if !isDigit {
					numEndIndex--
				}

				num := getNumFromState(line, numEndIndex+1, currentNumber)

				beginningIndexToCheck := max(numEndIndex-len(currentNumber), 0)
				lastIndexToCheck := min(numEndIndex+1, len(line)-1)
				// check before and after on this line
				if (beginningIndexToCheck >= 0 && isSymbol(line[beginningIndexToCheck])) || (lastIndexToCheck < len(line) && isSymbol(line[lastIndexToCheck])) {
					total += num
					currentNumber = []rune{}
					continue
				}

				for i := beginningIndexToCheck; i <= lastIndexToCheck; i++ {
					if y-1 >= 0 && isSymbol(schematic[y-1][i]) {
						// found it!
						total += num
						currentNumber = []rune{}
						continue
					}

					if y+1 < len(schematic) && isSymbol(schematic[y+1][i]) {
						// found it!
						total += num
						currentNumber = []rune{}
						continue
					}
				}

				currentNumber = []rune{}
			}
		}
	}

	fmt.Println(total)
}

func getNumFromState(line string, x int, state []rune) int64 {
	num, err := strconv.ParseInt(line[x-len(state):x], 0, 16)

	if err != nil {
		log.Fatal("couldn't parse a number", err)
	}

	return num
}

func isSymbol(char byte) bool {
	if char == 47 {
		return true
	}

	return char < 46 || char > 57
}
