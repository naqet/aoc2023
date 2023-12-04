package days

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/naqet/aoc2023/days"
)

func Logic() {
	input := days.ProcessInput("3");

    result := part1(input);
    fmt.Println("Part 1: ", result);

    result = part2(input);
    fmt.Println("Part 2: ", result)
}

func part1(input []string) int {
    var result int;
    for rowIdx, line := range input {
        // -1 means it doesn't, every other is the start index
        startNumberIdx := -1;
        for i, char := range line {
            if unicode.IsDigit(char) {
                if startNumberIdx == -1 {
                    startNumberIdx = i;
                }

                if i == len(line) - 1 {
                    number := line[startNumberIdx:]
                    if parsed, err := strconv.Atoi(number); err == nil {
                        if check(input, rowIdx, startNumberIdx, i) {
                            result += parsed;
                        }
                    }

                    startNumberIdx = -1;
                }
            } else if startNumberIdx != -1 {
                number := line[startNumberIdx:i]
                if parsed, err := strconv.Atoi(number); err == nil {
                    if check(input, rowIdx, startNumberIdx, i) {
                        result += parsed;
                    }
                }

                startNumberIdx = -1;
            }
        }
    }
    return result;
}

func check(input []string, rowIdx int, start int, end int) bool {
    row := input[rowIdx];
    lastRowIdx := len(row) - 1;
    var prev string;
    var next string;

    // If first
    if rowIdx != 0 {
        prev = input[rowIdx - 1];
    }

    // If last
    if rowIdx != len(input) - 1 {
        next = input[rowIdx + 1];
    }

    allowed := false;

    // Prev || Next
    if (start != 0 && isAllowed(string(row[start - 1]))) || (end != lastRowIdx && isAllowed(string(row[end]))) {
        return true;
    }

    if prev != "" {
        var startIdx int;
        if start == 0 {
            startIdx = start;
        } else {
            startIdx = start - 1;
        }

        for i := startIdx; i <= end; i++ {
            char := string(prev[i]);
             if isAllowed(char) {
                allowed = true;
            }
        }
    }

    if allowed {
        return allowed;
    }

    if next != "" {
        var startIdx int;
        if start == 0 {
            startIdx = start;
        } else {
            startIdx = start - 1;
        }

        for i := startIdx; i <= end; i++ {
            char := string(next[i]);
            if isAllowed(char) {
                allowed = true
            }
        }
    }

    return allowed;
}

func isAllowed(char string) bool {
    _, err := strconv.Atoi(char);
    return err != nil && char != ".";
}


func part2(input []string) int {
    var result int;
    for rowIdx, line := range input {
        for i, char := range strings.Split(line, "") {
            if char == "*" {
                numbers := getSiblingNums(line, i);
                if rowIdx != 0 {
                    rowNums := getRowNums(input[rowIdx - 1], i);
                    numbers = append(numbers, rowNums...);
                }

                if rowIdx != len(input) - 1 {
                    rowNums := getRowNums(input[rowIdx + 1], i);
                    numbers = append(numbers, rowNums...);
                }

                if len(numbers) == 2 {
                    result += numbers[0] * numbers[1];
                }
            }
        }
    }
    return result;
}

func getSiblingNums(line string, i int) []int {
    numbers := make([]int, 0);

    // To the left
    if i != 0 {
        if num, err := intAtEnd(line[:i]); err == nil {
            numbers = append(numbers, num);
        }
    }

    // To the right
    if i != len(line) + 1 {
        if num, err := intAtBeginning(line[i + 1:]); err == nil {
            numbers = append(numbers, num);
        }
    }

    return numbers;
}

func getRowNums(line string, i int) []int {
    nums := make([]int, 0);
    var isMiddle bool

    if _, err := strconv.Atoi(string(line[i])); err == nil {
        isMiddle = true
    }

    if isMiddle {
        var lmid bool;
        var midr bool;

        if i != 0 {
            _, err := strconv.Atoi(line[i-1:i+1]);

            if err == nil {
                lmid = true;
            }
        }

        if i != len(line) - 1 {
            _, err := strconv.Atoi(line[i:i+2]);

            if err == nil {
                midr = true;
            }
        }

        if lmid && midr {
            lIdx := i - 1;
            rIdx := i + 1;
            for idx := i; idx > -1; idx-- {
                char := line[idx];
                if unicode.IsDigit(rune(char)) {
                    lIdx = idx;
                } else {
                    break;
                }
            }

            for idx := i; idx < len(line); idx++ {
                char := line[idx];
                if unicode.IsDigit(rune(char)) {
                    rIdx = idx;
                } else {
                    break;
                }
            }

            midNum := line[lIdx:rIdx + 1];

            if parsed, err := strconv.Atoi(midNum); err == nil {
                nums = append(nums, parsed);
            }
        } else if lmid {
            if left, err := intAtEnd(line[:i + 1]); err == nil {
                nums = append(nums, left);
            }
        } else if midr {
            if right, err := intAtBeginning(line[i:]); err == nil {
                nums = append(nums, right);
            }
        } else {
            if midNum, err := strconv.Atoi(string(line[i])); err == nil {
                nums = append(nums, midNum);
            }
        }
    } else {
        if left, err := intAtEnd(line[:i]); err == nil {
            nums = append(nums, left);
        }
        if right, err := intAtBeginning(line[i + 1:]); err == nil {
            nums = append(nums, right);
        }
    }
    return nums;
}

func intAtEnd(word string) (int, error) {
    lastIndex := strings.LastIndexFunc(word, func(r rune) bool {return !unicode.IsDigit(r)});

    if lastIndex == -1 {
        parsed, err := strconv.Atoi(word);
        return parsed, err;
    }

    possibleInt := word[lastIndex + 1:];

    parsed, err := strconv.Atoi(possibleInt)
    return parsed, err;
}

func intAtBeginning(word string) (int, error) {
    idx := strings.IndexFunc(word, func(r rune) bool {return !unicode.IsDigit(r)});

    if idx == -1 {
        parsed, err := strconv.Atoi(word);
        return parsed, err;
    }

    possibleInt := word[:idx];

    parsed, err := strconv.Atoi(possibleInt)
    return parsed, err;
}
