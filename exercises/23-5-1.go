package exercises

import (
	"fmt"
	"strings"

	"8thday.dev/aog/utils"
)

func getLocation(val int64, farm *map[string]NutrientMapping) int64 {
	location := int64(-1)

	sourceNutrient := "seed"
	nextSourceValue := val

	for location < 0 {
		nutrientMapping := (*farm)[sourceNutrient]

		nextSourceValue = nutrientMapping.getDestinationValue(nextSourceValue)
		sourceNutrient = nutrientMapping.dest

		if sourceNutrient == "location" {
			location = nextSourceValue
		}
	}

	return location
}

func Year23Day5Exercise1() {
	input := utils.LineReaderString("./23-5.txt")

	sections := strings.Split(strings.Join(input, "\n"), "\n\n")

	seeds := utils.MapStringToInt64(strings.Fields(strings.Replace(sections[0], "seeds:", "", 1)))

	maps := sections[1:]

	farm := map[string]NutrientMapping{}

	for _, m := range maps {
		nutrientMapping := *NewNutrientMapping(strings.Split(strings.Trim(m, " \n"), "\n"))

		farm[nutrientMapping.source] = nutrientMapping
	}

	lowestLocation := int64(0)
	for i, s := range seeds {
		thisLocation := getLocation(s, &farm)

		if i == 0 || thisLocation < lowestLocation {
			lowestLocation = thisLocation
		}
	}

	fmt.Println(lowestLocation)
}
