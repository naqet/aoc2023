package days

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/naqet/aoc2023/days"
)

func Logic() {
	input := days.ProcessInput("1");
    result := part1(input);

    fmt.Println("Part 1: ", result);

    result = part2(input);

    fmt.Println("Part 2: ", result)
}

func part1(input []string) int {
    result := 0
    for _, line := range input {
        numbers := make([]int, 0)
        for _, char := range strings.Split(line, "") {
            if parsed, err := strconv.Atoi(char); err == nil {
                numbers = append(numbers, parsed);
                continue
            }
        }
        first := numbers[0]
        last := numbers[len(numbers)-1]
        result += 10*first + last
    }
    return result
}

func part2(input []string) int {
    result := 0
    digits := []string{
        "one", "two", "three", "four", "five",
        "six", "seven", "eight", "nine",
    }

    for _, line := range input {
        numbers := make([]int, 0)
        for i, char := range strings.Split(line, "") {
            if parsed, err := strconv.Atoi(char); err == nil {
                numbers = append(numbers, parsed);
                continue
            }
            for j, digit := range digits {
                if strings.HasPrefix(string(line[i:]), digit) {
                    numbers = append(numbers, j+1)
                }
            }
        }
        first := numbers[0]
        last := numbers[len(numbers)-1]
        result += 10*first + last
    }
    return result

}
