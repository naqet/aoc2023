package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/naqet/aoc2023/days"
)

func Logic() {
	input := days.ProcessInput("5")

	result := part1(input)
	fmt.Println("Part 1: ", result)

	result = part2(input)
	fmt.Println("Part 2: ", result)
}

func part1(input []string) int {
	lastEmpty := 0
	matrix := make([][]string, 0)
    seeds := make([]int, 0);
	for i, line := range input {
		if i == 0 {
            seeds = getSeeds(line);
		}

		if line == "" {
			if lastEmpty != 0 {
				matrix = append(matrix, input[lastEmpty+2:i])
			}

			lastEmpty = i
		}

        if i == len(input) - 1 {
            matrix = append(matrix, input[lastEmpty+2:])
        }
	}

	for _, stage := range matrix {
        processed := make([]int, 0);

		for _, stringNums := range stage {
            nums := strings.Split(stringNums, " ");
			destination, err := strconv.Atoi(nums[0])

			if err != nil {
				log.Fatal(err)
			}

			source, err := strconv.Atoi(nums[1])

			if err != nil {
				log.Fatal(err)
			}

			length, err := strconv.Atoi(nums[2])

			if err != nil {
				log.Fatal(err)
			}
            
            for i, v := range seeds {
                if (v > source || v == source) && (v < source + length || v == source + length) {
                    res := v - source;
                    num := destination + res;
                    if !isInArray(processed, v) {
                        seeds[i] = num;
                        processed = append(processed, num)
                    }
                }
            }
		}
	}

    var value int;

    for i, v := range seeds {
        if i == 0 || value > v {
            value = v
        }
        
    }
	return value;
}

func part2(input []string) int {
    return 0;
}

func getSeeds(line string) []int {
	stringSeeds := strings.Split(line[7:], " ")
    seeds := make([]int, 0);

    for _, v := range stringSeeds {
        parsed, err := strconv.Atoi(v);

        if err != nil {
            log.Fatal(err);
        }
        seeds = append(seeds, parsed);
    }

	return seeds;
}

func isInArray(arr []int, target int) bool {
    found := false;
    for _, v := range arr {
        if v == target {
            found = true;
            break;
        }
    }

    return found;
}
