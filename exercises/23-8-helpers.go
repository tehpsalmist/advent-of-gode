package exercises

import "strings"

func travelGraph(instructions string, graph map[string]map[rune]string, startingNode string) int64 {
	total := int64(0)
	currentNode := startingNode

	// coincidentally, this check also works for part 1
	for !strings.HasSuffix(currentNode, "Z") {
		for _, direction := range instructions {
			total++
			currentNode = graph[currentNode][direction]

			if strings.HasSuffix(currentNode, "Z") {
				break
			}
		}
	}

	return total
}
