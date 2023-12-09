package exercises

import (
	"fmt"

	"8thday.dev/aog/utils"
)

func Year23Day1Exercise1() {
	total := 0

	utils.LineReaderBytesIterator("./inputs/23-1.txt", func(line []byte) {
		for i := 0; i < len(line); i++ {
			if 48 <= line[i] && line[i] <= 57 {
				total = total + ((int(line[i]) - 48) * 10)
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if 48 <= line[i] && line[i] <= 57 {
				total = total + (int(line[i]) - 48)
				break
			}
		}
	})

	fmt.Println(total)
}
