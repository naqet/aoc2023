package main

import (
	"log"
	"os"
	"strconv"

	day1 "github.com/naqet/aoc2023/days/1"
	day2 "github.com/naqet/aoc2023/days/2"
	day3 "github.com/naqet/aoc2023/days/3"
	day4 "github.com/naqet/aoc2023/days/4"
	day5 "github.com/naqet/aoc2023/days/5"
)

func main() {
    funcs := map[int]func(){
        1: day1.Logic,
        2: day2.Logic,
        3: day3.Logic,
        4: day4.Logic,
        5: day5.Logic,
    }

	if len(os.Args) != 2 {
        log.Fatal("Incorrect number of args");
	}

    day := os.Args[1];

    if day, err := strconv.Atoi(day); err == nil {
        if day < 0 || day > len(funcs) {
            log.Fatal("Incorrect day name");
        }

        exec := funcs[day];
        exec();
    } else {
        log.Fatal("Day arg is not a number");
    }
}
