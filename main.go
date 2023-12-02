package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"github.com/naqet/aoc2023/days"
)

type inputArgs struct {
	day int
}

func main() {
    input, err := handleInput();

    if err != nil {
        log.Fatal(err)
    }

    dayFuncs := days.GetDayFunctions();

    if input.day < 0 || input.day > len(dayFuncs) {
        log.Fatal("Day arg is larger than number of available day functions");
    }

    funcToExec := dayFuncs[input.day];

    funcToExec();
}

func handleInput() (inputArgs, error) {
	if len(os.Args) != 2 {
        return inputArgs{}, errors.New("Incorrect number of arguments");
	}

    day := os.Args[1];

    if day, err := strconv.Atoi(day); err == nil {
        return inputArgs{day}, nil
    }

    return inputArgs{}, errors.New("Day arg is not a number");
}
