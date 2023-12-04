package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/naqet/aoc2023/days"
)

type balls struct {
    red int
    blue int
    green int
}

func Logic() {
	input := days.ProcessInput("2");
    result := part1(input);

    fmt.Println("Part 1: ", result);

    result = part2(input);

    fmt.Println("Part 2: ", result)
}

func part1(input []string) int {
    var result int;
    for _, line := range input {
        id := getId(line);
        subsets := getSubsets(line);
        isEnough := isEnoughBalls(subsets);

        if isEnough {
            if parsed, err := strconv.Atoi(id); err == nil {
                result += parsed;
            }
        }
    }

    return result;
}

func part2(input []string) int {
    var result int;
    for _, line := range input {
        subsets := getSubsets(line);
        ballCount := getFewestCount(subsets);

        var localResult int;

        for _, v := range ballCount {
            if localResult == 0 {
                localResult += v;
            } else {
                localResult *= v;
            }
        }

        result += localResult;
    }

    return result;
}

func getId(line string) string {
    firstSpace := strings.Index(line, " ");
    firstSemicolon := strings.Index(line, ":");
    return line[firstSpace + 1:firstSemicolon];
}

func getSubsets(line string) []string {
    firstSemicolon := strings.Index(line, ":");
    lineWithoutId := line[firstSemicolon + 2:];
    subsets := strings.Split(lineWithoutId, "; ")
    return subsets;
}

func isEnoughBalls(subsets []string) bool {
    enough := true;

    ballLimit := map[string]int {
        "red": 12,
        "green": 13,
        "blue": 14,
    }

    out: for _, set := range subsets {
        for key := range ballLimit {
            count := getBallCount(key, set);
            if count > ballLimit[key] {
               enough = false;
               break out; 
            }
        }
    }

    return enough;
}

func getFewestCount(subsets []string) []int {
    ballCount := map[string]int{};

    colors := []string{
        "red",
        "green",
        "blue",
    }

    for _, set := range subsets {
        for _, color := range colors {
            count := getBallCount(color, set);
            if ballCount[color] < count {
                ballCount[color] = count;
            }
        }
    }
    var counts []int;

    for _, v := range ballCount {
        counts = append(counts, v);
    }
    return counts;
}

func getBallCount(color string, input string) int {
    instructions := strings.Split(input, ", ")

    var a string;

    for _, stmt := range instructions {
        if strings.Contains(stmt, color) {
            a = stmt;
            break;
        }
    }

    number := strings.Split(a, " ")[0];

    if parsed, err := strconv.Atoi(number); err == nil {
        return parsed;
    }

    return 0;
}
