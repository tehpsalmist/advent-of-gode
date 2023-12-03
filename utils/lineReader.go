package utils

import (
	"bufio"
	"log"
	"os"
)

func LineReaderBytes(filePath string) [][]byte {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	text := [][]byte{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text = append(text, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}

func LineReaderString(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	text := []string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}

func LineReaderBytesIterator(filePath string, consumer func(line []byte)) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		consumer(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func LineReaderStringIterator(filePath string, consumer func(line string)) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		consumer(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
