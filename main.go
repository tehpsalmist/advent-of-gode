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
	case "23-2-1":
		exercises.Year23Day2Exercise1()
	case "23-2-2":
		exercises.Year23Day2Exercise2()
	case "23-3-1":
		exercises.Year23Day3Exercise1()
	case "23-3-2":
		exercises.Year23Day3Exercise2()
	case "23-4-1":
		exercises.Year23Day4Exercise1()
	case "23-4-2":
		exercises.Year23Day4Exercise2()
	case "23-5-1":
		exercises.Year23Day5Exercise1()
	case "23-5-2":
		exercises.Year23Day5Exercise2()
	case "23-6-1":
		exercises.Year23Day6Exercise1()
	case "23-6-2":
		exercises.Year23Day6Exercise2()
	case "23-7-1":
		exercises.Year23Day7Exercise1()
	case "23-7-2":
		exercises.Year23Day7Exercise2()
	case "23-8-1":
		exercises.Year23Day8Exercise1()
	case "23-8-2":
		exercises.Year23Day8Exercise2()
	case "23-9-1":
		exercises.Year23Day9Exercise1()
	case "23-9-2":
		exercises.Year23Day9Exercise2()
	case "23-10-1":
		exercises.Year23Day10Exercise1()
	case "23-10-2":
		exercises.Year23Day10Exercise2()
	case "23-11-1":
		exercises.Year23Day11Exercise1()
	case "23-11-2":
		exercises.Year23Day11Exercise2()
	case "23-12-1":
		exercises.Year23Day12Exercise1()
	case "23-12-2":
		exercises.Year23Day12Exercise2()
	case "23-13-1":
		exercises.Year23Day13Exercise1()
	case "23-13-2":
		exercises.Year23Day13Exercise2()
	case "23-14-1":
		exercises.Year23Day14Exercise1()
	case "23-14-2":
		exercises.Year23Day14Exercise2()
	case "23-15-1":
		exercises.Year23Day15Exercise1()
	case "23-15-2":
		exercises.Year23Day15Exercise2()
	case "23-16-1":
		exercises.Year23Day16Exercise1()
	case "23-16-2":
		exercises.Year23Day16Exercise2()
	case "23-17-1":
		exercises.Year23Day17Exercise1()
	case "23-17-2":
		exercises.Year23Day17Exercise2()
	case "23-18-1":
		exercises.Year23Day18Exercise1()
	case "23-18-2":
		exercises.Year23Day18Exercise2()
	case "23-19-1":
		exercises.Year23Day19Exercise1()
	case "23-19-2":
		exercises.Year23Day19Exercise2()
	case "23-20-1":
		exercises.Year23Day20Exercise1()
	case "23-20-2":
		exercises.Year23Day20Exercise2()
	case "23-21-1":
		exercises.Year23Day21Exercise1()
	case "23-21-2":
		exercises.Year23Day21Exercise2()
	case "23-22-1":
		exercises.Year23Day22Exercise1()
	case "23-22-2":
		exercises.Year23Day22Exercise2()
	case "23-23-1":
		exercises.Year23Day23Exercise1()
	case "23-23-2":
		exercises.Year23Day23Exercise2()
	case "23-24-1":
		exercises.Year23Day24Exercise1()
	case "23-24-2":
		exercises.Year23Day24Exercise2()
	case "23-25-1":
		exercises.Year23Day25Exercise1()
	case "23-25-2":
		exercises.Year23Day25Exercise2()
	default:
		exerciseData := strings.Split(os.Args[1], "-")

		fmt.Printf("No code found for Day %s, Exercise %s in Year 20%s", exerciseData[1], exerciseData[2], exerciseData[0])
	}
}
