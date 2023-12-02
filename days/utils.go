package days

import (
	"bufio"
	"log"
	"os"
)

func ProcessInput(day string) []string {
	file, err := os.Open("days/" + day + "/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
