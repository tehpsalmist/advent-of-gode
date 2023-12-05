package exercises

import (
	"fmt"
	"strings"

	"8thday.dev/aog/utils"
)

var shouldLog = true

type MapRange struct {
	source int64
	dest   int64
	cap    int64
	offset int64
}

func NewMapRange(input string) *MapRange {
	mr := &MapRange{}

	values := utils.MapStringToInt64(strings.Fields(input))

	mr.source = values[1]
	mr.dest = values[0]
	mr.cap = mr.source + values[2]
	mr.offset = mr.dest - mr.source

	return mr
}

func (mr *MapRange) checkValue(val int64) (int64, bool) {

	if mr.source <= val && val < mr.cap {
		return val + mr.offset, true
	}

	return -1, false
}

type NutrientMapping struct {
	ranges []MapRange
	source string
	dest   string
}

func NewNutrientMapping(input []string) *NutrientMapping {
	nm := &NutrientMapping{}

	nutrients := strings.Split(strings.Split(input[0], " ")[0], "-")

	nm.source = nutrients[0]
	nm.dest = nutrients[2]

	mappings := input[1:]

	nm.ranges = make([]MapRange, len(mappings))

	for i, m := range mappings {
		nm.ranges[i] = *NewMapRange(m)
	}

	return nm
}

func (nm *NutrientMapping) getDestinationValue(val int64) int64 {

	for _, r := range nm.ranges {
		finalValue, ok := r.checkValue(val)

		if ok {
			return finalValue
		}
	}

	return val
}

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
