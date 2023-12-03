package exercises

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"8thday.dev/aog/utils"
)

var colorConstraints = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Year23Day2Exercise1() {
	total := 0

	utils.LineReaderStringIterator("./23-2.txt", func(line string) {
		game := strings.Split(line, ": ")

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

				if value > int64(colorConstraints[color]) {
					return
				}
			}
		}

		idVal, err := strconv.ParseInt(strings.Replace(game[0], "Game ", "", 1), 0, 16)

		if err != nil {
			log.Fatal(err)
		}

		total += int(idVal)
	})

	fmt.Println(total)
}
