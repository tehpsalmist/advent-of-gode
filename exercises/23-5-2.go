package exercises

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day5Exercise2() {
	bestSeed := int64(-1)

	input := utils.LineReaderString("./inputs/23-5-test.txt")

	sections := strings.Split(strings.Join(input, "\n"), "\n\n")

	seedInput := utils.MapStringToInt64(strings.Fields(strings.Replace(sections[0], "seeds:", "", 1)))

	seeds := [][]int64{}
	for i := 0; i < len(seedInput)-2; i += 2 {
		seedStart := seedInput[i]
		seedEnd := seedStart + (seedInput[i+1] - 1)
		seeds = append(seeds, []int64{seedStart, seedEnd})
	}

	maps := sections[1:]

	farm := map[string]NutrientMapping{}

	for _, m := range maps {
		nutrientMapping := *NewNutrientMapping(strings.Split(strings.Trim(m, " \n"), "\n"))

		farm[nutrientMapping.dest] = nutrientMapping
	}

	locationMapping := farm["location"]
	for _, lr := range locationMapping.ranges {
		parameters := [][]int64{{lr.source, lr.sourceCap, lr.dest}}

		for _, nutrient := range []string{"humidity", "temperature", "light", "water", "fertilizer", "soil"} {
			nextMapping := farm[nutrient]

			newParameters := [][]int64{}
			for _, param := range parameters {
				lowerLimit := param[0]
				upperLimit := param[1]
				// find the starting point and ending point, translate to source values, insert into
				intermediateParams := [][]int64{}
				for _, r := range nextMapping.ranges {
					sortMarker := int64(-1)
					newLow := int64(-1)
					newEnd := int64(-1)

					// set new low
					if lowerLimit <= r.dest && r.dest <= upperLimit {
						newLow = r.source // translate to source
						sortMarker = r.dest
					} else if r.dest < lowerLimit && lowerLimit <= r.destCap {
						newLow = lowerLimit - r.offset // translate to source
						sortMarker = lowerLimit
					} else {
						// these ranges don't overlap, next iteration!
						continue
					}

					if lowerLimit <= r.destCap && r.destCap <= upperLimit {
						newEnd = r.sourceCap // translate to source
					} else if r.dest < upperLimit && upperLimit <= r.destCap {
						newEnd = upperLimit - r.offset // translate to source
					}

					intermediateParams = append(intermediateParams, []int64{newLow, newEnd, sortMarker})
				}

				sort.Slice(intermediateParams, func(i, j int) bool { return intermediateParams[i][2] < intermediateParams[j][2] })
				newParameters = append(newParameters, intermediateParams...)
			}

			parameters = newParameters

			if nutrient == "soil" {
				foundBest, ok := findBestSeed(parameters, seeds)

				if ok {
					bestSeed = foundBest
					break
				}
			}
		}
		if bestSeed >= 0 {
			break
		}
	}

	fmt.Println(bestSeed)
}

func findBestSeed(parameters, seeds [][]int64) (int64, bool) {
	lowestSeed := int64(math.MaxInt64)

	for _, r := range parameters {
		lowerLimit := r[0]
		upperLimit := r[1]

		for _, s := range seeds {
			start := s[0]
			end := s[1]

			fmt.Println("start", start)
			fmt.Println("end", end)

			if lowerLimit == start {
				fmt.Println("found start")
				return start, true
			}

			if lowerLimit < start && start <= upperLimit {
				lowestSeed = min(lowestSeed, start)
				fmt.Println("incremented lowest seed")
				continue
			}

			if lowerLimit <= end && end <= upperLimit {
				fmt.Println("found lower limit")
				return lowerLimit, true
			}
		}
	}

	fmt.Println("found lowestSeed")
	return lowestSeed, false
}
