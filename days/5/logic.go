package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type seedPair struct {
    start int
    end int
}

type rangeDetails struct {
    destination int
    source int
    length int
}

func Logic() {
	seeds, blocks := getInput();

	result := part1(seeds, blocks)
	fmt.Println("Part 1: ", result)

	result = part2(seeds, blocks);
	fmt.Println("Part 2: ", result)

}

func part1(input []int, blocks [][]string) int {
    seeds := input;
    
    for _, v := range blocks {
        ranges := make([]rangeDetails, 0);

        for _, details := range v {
            nums := strings.Fields(details);
            destination, _ := strconv.Atoi(nums[0]);
            source, _ := strconv.Atoi(nums[1]);
            length, _ := strconv.Atoi(nums[2]);
            
            ranges = append(ranges, rangeDetails{destination, source, length});
        }

        newSeeds := make([]int, 0);

        for _, s := range seeds {
            inRange := false;
            for _, r := range ranges {
                if r.source <= s && s < r.source + r.length {
                    newSeeds = append(newSeeds, s - r.source + r.destination)
                    inRange = true;
                    break;
                }
            }

            if !inRange {
                newSeeds = append(newSeeds, s);
            }
        }

        seeds = newSeeds;
    }

    lowest := seeds[0]

    for _, s := range seeds {
        if lowest > s {
            lowest = s
        }
    }

    return lowest;
}

func part2(input []int, blocks [][]string) int {
    seeds := make([]seedPair, 0);

    for i, v := range input {
        if i % 2 == 0 {
            seeds = append(seeds, seedPair{v, v + input[i + 1]})
        }
    }
    for _, v := range blocks {
        ranges := make([]rangeDetails, 0);

        for _, details := range v {
            nums := strings.Fields(details);
            destination, _ := strconv.Atoi(nums[0]);
            source, _ := strconv.Atoi(nums[1]);
            length, _ := strconv.Atoi(nums[2]);
            
            ranges = append(ranges, rangeDetails{destination, source, length});
        }

        newSeeds := make([]seedPair, 0);

        for len(seeds) > 0 {
            s := seeds[len(seeds)-1];
            seeds = seeds[:len(seeds)-1];

            overlap := false;

            for _, r := range ranges {
                os := max(s.start, r.source);
                oe := min(s.end, r.source + r.length);

                if os < oe {
                    overlap = true;
                    newSeeds = append(newSeeds, seedPair{os - r.source + r.destination, oe - r.source + r.destination})

                    if os > s.start {
                        seeds = append(seeds, seedPair{s.start, os});
                    }

                    if oe < s.end {
                        seeds = append(seeds, seedPair{oe, s.end});
                    }
                    break;
                }
            }

            if !overlap {
                newSeeds = append(newSeeds, s);
            }
        }

        seeds = newSeeds;
    }

    lowest := seeds[0].start;

    for _, s := range seeds {
        if lowest > s.start {
            lowest = s.start
        }
    }

    return lowest;
}

func getInput() ([]int, [][]string) {
	file, err := os.Open("days/5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var blocks [][]string;
    var seeds []int;

	scanner := bufio.NewScanner(file)

    scanner.Scan();

    text := scanner.Text();

    if strings.Contains(text, "seeds") {
        for _, v := range strings.Fields(text[7:]) {
            num, _ := strconv.Atoi(v);
            seeds = append(seeds, num);
        }
    }
    scanner.Scan();

    var newBlock []string;

	for scanner.Scan() {
        text := scanner.Text();
        if strings.Contains(text, "map:") {
            continue;
        }

        if text == "" {
            blocks = append(blocks, newBlock);
            newBlock = []string{};
        } else {
            newBlock = append(newBlock, text);
        }
	}

    blocks = append(blocks, newBlock);

	return seeds, blocks;
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
