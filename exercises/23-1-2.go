package exercises

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"8thday.dev/aog/utils"
)

func Year23Day1Exercise2() {
	numberTrees := *loadNumberTree()
	treeFromStart := numberTrees[0]

	treeFromEnd := numberTrees[1]

	total := 0

	utils.LineReaderStringIterator("./23-1.txt", func(line string) {
		firstValue := findValueFromLine(line, treeFromStart, false)
		lastValue := findValueFromLine(line, treeFromEnd, true)

		total += firstValue * 10
		total += lastValue
	})

	fmt.Println(total)
}

func findValueFromLine(line string, initialTree map[string]any, backwards bool) int {
	/* Silly Boilerplate for being able to iterate backwards */
	i := 0
	limit := len(line)
	advanceFunc := func(currVal int) int {
		return currVal + 1
	}
	compareFunc := func(a int, b int) bool {
		return a < b
	}

	if backwards {
		i = len(line) - 1
		limit = 0
		advanceFunc = func(currVal int) int {
			return currVal - 1
		}
		compareFunc = func(a int, b int) bool {
			return a >= b
		}
	}
	/* End Silly Boilerplate */

	traversingTree := initialTree
	currentWord := []byte{}

	for ; compareFunc(i, limit); i = advanceFunc(i) {
		character := line[i]

		// this byte is an integer! Yay! Convert to int and return!
		if 48 <= character && character <= 57 {
			foundValue, _ := strconv.Atoi(string(character))

			return foundValue
		}

		// save state so we can backtrack later if necessary
		currentWord = append(currentWord, character)

		nextBranch, traversingTreeOk := traversingTree[string(character)].(map[string]any)

		// found a character we were expecting in the word tree!
		if traversingTreeOk {
			// advance the tree down this branch for the next iteration
			traversingTree = nextBranch

			// it's possible this is actually the final letter in a word,
			// if so, it will hold its value for us to access
			found, ok := traversingTree["value"].(float64)

			// awesome, we found the value, convert and return!
			if ok {
				return int(found)
			}
		} else if len(currentWord) > 1 {
			// we hit a dead end with the current branch we were going down: this character wasn't a match
			// but since we have some characters stored in our current word,
			// they might be a successful path on their own, so let's investigate!

			// our candidate path is all but the first letter, since we know starting there will dead end
			candidates := currentWord[1:]
			// reset word state and go back to the top of the tree
			currentWord = []byte{}
			traversingTree = initialTree

			// iterate through the letters to see if it yields a path
			for i := 0; i < len(candidates); i++ {
				nextBranch, traversingTreeOk := traversingTree[string(candidates[i])].(map[string]any)

				// found stuff
				if traversingTreeOk {
					traversingTree = nextBranch
					currentWord = append(currentWord, candidates[i])
				} else {
					currentWord = []byte{}
					traversingTree = initialTree
				}
			}
		} else {
			// current state and character all yielded nothing, reset everything for the next one
			currentWord = []byte{}
			traversingTree = initialTree
		}
	}

	return 0
}

func loadNumberTree() *[]map[string]any {
	f, err := os.ReadFile("./number-tree.json")

	if err != nil {
		log.Fatal(err)
	}

	numberTree := &[]map[string]any{}

	err = json.Unmarshal(f, numberTree)

	if err != nil {
		log.Fatal(err)
	}

	return numberTree
}
