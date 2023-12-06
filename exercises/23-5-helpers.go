package exercises

import (
	"fmt"
	"sort"
	"strings"

	"8thday.dev/aog/utils"
)

type MapRange struct {
	source    int64
	dest      int64
	sourceCap int64
	destCap   int64
	offset    int64
}

func NewMapRange(input string) *MapRange {
	mr := &MapRange{}

	values := utils.MapStringToInt64(strings.Fields(input))

	mr.source = values[1]
	mr.dest = values[0]
	mr.sourceCap = mr.source + (values[2] - 1)
	mr.offset = mr.dest - mr.source
	mr.destCap = mr.dest + (values[2] - 1)

	return mr
}

func (mr *MapRange) checkValue(val int64) (int64, bool) {

	if mr.source <= val && val <= mr.sourceCap {
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

	definedRanges := make([]MapRange, len(mappings))

	for i, m := range mappings {
		definedRanges[i] = *NewMapRange(m)
	}

	sort.Slice(definedRanges, func(i, j int) bool { return definedRanges[i].dest < definedRanges[j].dest })

	gaps := []MapRange{}
	for i, m := range definedRanges {
		if i == len(definedRanges)-1 {
			// cover our butts with the upper range gap
			gaps = append(gaps, *NewMapRange(fmt.Sprint(m.destCap+1, m.destCap+1, 9000000000)))
			continue
		}

		if i == 0 && m.dest != 0 {
			gaps = append(gaps, *NewMapRange((fmt.Sprint(0, 0, m.dest))))
		}

		difference := definedRanges[i+1].dest - m.destCap
		if difference != 1 {
			gaps = append(gaps, *NewMapRange(fmt.Sprint(m.destCap+1, m.destCap+1, difference-1)))
		}
	}

	nm.ranges = append(definedRanges, gaps...)
	sort.Slice(nm.ranges, func(i, j int) bool { return nm.ranges[i].dest < nm.ranges[j].dest })

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
