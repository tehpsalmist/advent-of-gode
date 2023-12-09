package exercises

import (
	"fmt"
	"log"
	"strconv"

	"8thday.dev/aog/utils"
)

func Year23Day3Exercise2() {
	total := int64(0)

	schematic := utils.LineReaderString("./inputs/23-3.txt")

	for y, line := range schematic {
		for x, char := range line {
			if isGear(char) {
				locations := [][]int{}

				beginningIndexToCheck := max(x-1, 0)
				lastIndexToCheck := min(x+1, len(line)-1)

				// check before and after on this line
				if byteIsDigit(line[beginningIndexToCheck]) {
					locations = append(locations, []int{beginningIndexToCheck, y})
				}

				if byteIsDigit(line[lastIndexToCheck]) {
					locations = append(locations, []int{lastIndexToCheck, y})
				}

				if y-1 >= 0 {
					justFound := false
					for i := beginningIndexToCheck; i <= lastIndexToCheck; i++ {
						isDigit := byteIsDigit(schematic[y-1][i])

						if isDigit && !justFound {
							locations = append(locations, []int{i, y - 1})
							justFound = true
						} else if !isDigit {
							justFound = false
						}
					}
				}

				if y+1 < len(schematic) {
					justFound := false
					for i := beginningIndexToCheck; i <= lastIndexToCheck; i++ {
						isDigit := byteIsDigit(schematic[y+1][i])

						if isDigit && !justFound {
							locations = append(locations, []int{i, y + 1})
							justFound = true
						} else if !isDigit {
							justFound = false
						}
					}
				}

				if len(locations) == 2 {
					val1 := getValueAtLocation(locations[0], schematic)
					val2 := getValueAtLocation(locations[1], schematic)

					total += val1 * val2
				}
			}
		}
	}

	fmt.Println(total)
}
func byteIsDigit(b byte) bool {
	return 47 < b && b < 58
}

func isGear(char rune) bool {
	if char == 42 {
		return true
	}

	return false
}

func getValueAtLocation(location []int, schematic []string) int64 {
	line := schematic[location[1]]
	x := location[0]

	valBytes := []byte{line[x]}

	for i := x - 1; i >= 0; i-- {
		if byteIsDigit(line[i]) {
			valBytes = append([]byte{line[i]}, valBytes...)
		} else {
			break
		}
	}

	for i := x + 1; i < len(line); i++ {
		if byteIsDigit(line[i]) {
			valBytes = append(valBytes, line[i])
		} else {
			break
		}
	}

	val, err := strconv.ParseInt(string(valBytes), 0, 16)

	if err != nil {
		log.Fatal("unable to parse number")
	}

	return val
}
