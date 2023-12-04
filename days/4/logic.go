package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/naqet/aoc2023/days"
)

func Logic() {
	input := days.ProcessInput("4");

    result := part1(input);
    fmt.Println("Part 1: ", result);

    result = part2(input);
    fmt.Println("Part 2: ", result)
}

func part1(input []string) int {
    var result int;
    for _, line := range input {
        result += getPoints(line);
    }
    return result;
}

func getPoints(line string) int {
    matches := getMatches(line);
    var points int;

    for i := 0; i < matches; i++ {
        if i == 0 {
            points += 1;
        } else {
            points = points * 2;
        }
    }

    return points;
}

func getMatches(line string) int {
    semiIdx := strings.Index(line, ":");
    withoutId := line[semiIdx + 1:];
    pointsLists := strings.Split(withoutId, "|");

    winning := strings.Split(pointsLists[0][1:len(pointsLists[0])- 1], " ");
    nums := strings.Split(pointsLists[1][1:], " ");

    var matches int;

    out: for _, num := range nums {
        if _, err := strconv.Atoi(num); err != nil {
            continue;
        }
        for _, win := range winning {
            if _, err := strconv.Atoi(win); err != nil {
                continue;
            }
            if num == win {
                matches += 1;
                continue out;
            }
        }
    }

    return matches;
}

func part2(input []string) int {
    var result int;
    for _, line := range input {
        result += getNext(line, input);
    }
    return result;
}

func getNext(line string, input []string) int {
    id, err := getId(line);
    if err != nil {
        log.Fatal(err);
    }
    result := 1;
    matches := getMatches(line);
    for idx := 1; idx <= matches; idx++ {
        if id + idx - 1 < len(input) {
            result += getNext(input[id + idx - 1], input);
        }
    }
    return result;
}

func getId(line string) (int, error) {
    group := strings.Split(line, ": ")[0];
    id := group[strings.LastIndex(group, " ") + 1:];

    parsed, err := strconv.Atoi(id);

    return parsed, err;
}
