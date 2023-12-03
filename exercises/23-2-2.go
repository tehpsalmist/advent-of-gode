package exercises

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day2Exercise2() {
	total := int64(0)

	utils.LineReaderStringIterator("./23-2.txt", func(line string) {
		game := strings.Split(line, ": ")

		redMax := int64(0)
		greenMax := int64(0)
		blueMax := int64(0)

		rounds := strings.Split(game[1], "; ")

		for _, round := range rounds {
			groups := strings.Split(round, ", ")

			for _, group := range groups {
				splitGroup := strings.Split(group, " ")

				color := splitGroup[1]
				value, err := strconv.ParseInt(splitGroup[0], 0, 16)

				if err != nil {
					log.Fatal(err)
				}

				switch color {
				case "red":
					redMax = max(value, redMax)
				case "green":
					greenMax = max(value, greenMax)
				case "blue":
					blueMax = max(value, blueMax)
				default:
					log.Fatal("bruh what even is this")
				}
			}
		}

		total += redMax * greenMax * blueMax
	})

	fmt.Println(total)
}
