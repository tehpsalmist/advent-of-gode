package main

import (
	"fmt"
	"os"
	"strings"

	"8thday.dev/aog/exercises"
)

func main() {
	switch os.Args[1] {
	case "23-1-1":
		exercises.Year23Day1Exercise1()
	case "23-1-2":
		exercises.Year23Day1Exercise2()
	default:
		exerciseData := strings.Split(os.Args[1], "-")

		fmt.Printf("No code found for Day %s, Exercise %s in Year 20%s", exerciseData[1], exerciseData[2], exerciseData[0])
	}
}
