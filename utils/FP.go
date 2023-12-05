package utils

import (
	"log"
	"strconv"
)

func MapStringToInt64(strs []string) []int64 {
	ret := make([]int64, len(strs))

	for i, s := range strs {
		intVal, err := strconv.ParseInt(s, 0, 64)

		if err != nil {
			log.Fatal("passed a non-string value to map func", s, err)
		}

		ret[i] = intVal
	}

	return ret
}
