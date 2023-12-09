package exercises

import (
	"fmt"
	"strings"

	"8thday.dev/aog/utils"
)

func Year23Day8Exercise1() {
	graph := map[string]map[rune]string{}

	instructions := ""

	startingNode := "AAA"

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

		graph[parts[0]] = map[rune]string{
			76: children[0],
			82: children[1],
		}

		i++
	})

	total := travelGraph(instructions, graph, startingNode)

	fmt.Println(total)
}
