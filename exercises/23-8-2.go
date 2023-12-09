package exercises

import (
	"fmt"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day8Exercise2() {
	graph := map[string]map[rune]string{}

	instructions := ""

	currentNodes := []string{}

	i := 0
	utils.LineReaderStringIterator("./inputs/23-8.txt", func(line string) {
		if i == 0 {
			instructions = line
			i++
			return
		}

		if i == 1 {
			i++
			return
		}

		parts := strings.Split(line, " = ")

		children := strings.Split(parts[1][1:len(parts[1])-1], ", ")

		nodeKey := parts[0]

		if nodeKey[2] == 65 {
			currentNodes = append(currentNodes, nodeKey)
		}

		graph[nodeKey] = map[rune]string{
			76: children[0],
			82: children[1],
		}

		i++
	})

	destinations := []int64{}
	for _, n := range currentNodes {
		destinations = append(destinations, travelGraph(instructions, graph, n))
	}

	// calculate lowest common denominator of all path destinations
	lcmProduct := int64(1)
	for _, d := range destinations {
		lcmProduct = lcm(d, lcmProduct)
	}

	fmt.Println(lcmProduct)
}

func gcd(a int64, b int64) int64 {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a int64, b int64) int64 {
	return (a * b) / gcd(a, b)
}
