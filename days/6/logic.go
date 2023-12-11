package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/naqet/aoc2023/days"
)

type race struct {
	time   int
	record int
}

func Logic() {
	input := days.ProcessInput("6")

	result := part1(input)
	fmt.Println("Part 1: ", result)

	result = part2(input)
	fmt.Println("Part 2: ", result)

}

func part1(input []string) int {
	races := getRaces(input)

	var result int
	for _, r := range races {
		var numberOfWins int
		for i := 1; i < r.time; i++ {
			points := i * (r.time - i)

			if points > r.record {
				numberOfWins++
			}
		}

		if result == 0 {
			result = numberOfWins
		} else {
			result *= numberOfWins
		}
	}

	return result
}

func part2(input []string) int {
	race := getRace(input)

	var numberOfWins int
	for i := 1; i < race.time; i++ {
		points := i * (race.time - i)

		if points > race.record {
			numberOfWins++
		}
	}

	return numberOfWins
}

func getRace(input []string) race {
	line := input[0]
	strInts := line[strings.Index(line, ":")+1:]
	strTime := strings.Join(strings.Fields(strInts), "")
	time, _ := strconv.Atoi(strTime)

	line = input[1]
	strInts = line[strings.Index(line, ":")+1:]
	strRecord := strings.Join(strings.Fields(strInts), "")
	record, _ := strconv.Atoi(strRecord)

	return race{time, record}
}

func getRaces(input []string) []race {
	line := input[0]
	strInts := line[strings.Index(line, ":")+1:]
	times := make([]int, 0)
	for _, v := range strings.Fields(strInts) {
		parsed, _ := strconv.Atoi(v)
		times = append(times, parsed)
	}

	line = input[1]
	strInts = line[strings.Index(line, ":")+1:]
	records := make([]int, 0)
	for _, v := range strings.Fields(strInts) {
		parsed, _ := strconv.Atoi(v)
		records = append(records, parsed)
	}

	races := make([]race, 0)

	for i := 0; i < len(times); i++ {
		races = append(races, race{times[i], records[i]})
	}
	return races
}
