package main

import (
	_ "bufio"
	"fmt"
	"github.com/DianaSun97/AdventGolang/common"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Part struct {
	Text string
	Irow int
	Icol int
}

func main() {
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalln(err)
	}

	rows := strings.Split(fileContent, "\n")
	symbols := parse(rows, regexp.MustCompile(`[^.0-9]`))
	nums := parse(rows, regexp.MustCompile(`\d+`))

	partOneResult := 0
	for _, n := range nums {
		if hasAdjacentSymbol(symbols, n) {
			numInt, err := strconv.Atoi(n.Text)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}
			partOneResult += numInt
		}
	}

	fmt.Println("Part One:", partOneResult)

	gears := parse(rows, regexp.MustCompile(`\*`))
	numbers := parse(rows, regexp.MustCompile(`\d+`))

	partTwoResult := 0
	for _, g := range gears {
		neighbours := getAdjacentNumbers(numbers, []Part{g})

		if len(neighbours) == 2 {
			num1, err := strconv.Atoi(neighbours[0].Text)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}

			num2, err := strconv.Atoi(neighbours[1].Text)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}

			partTwoResult += num1 * num2
		}
	}

	fmt.Println("Part Two:", partTwoResult)
}

func parse(rows []string, rx *regexp.Regexp) []Part {
	var result []Part
	for irow := 0; irow < len(rows); irow++ {
		matches := rx.FindAllStringIndex(rows[irow], -1)
		for _, match := range matches {
			result = append(result, Part{
				Text: rows[irow][match[0]:match[1]],
				Irow: irow,
				Icol: match[0],
			})
		}
	}
	return result
}

func hasAdjacentSymbol(symbols []Part, num Part) bool {
	for _, s := range symbols {
		if isAdjacent(s, num) {
			return true
		}
	}
	return false
}

func isAdjacent(p1, p2 Part) bool {
	return abs(p2.Irow-p1.Irow) <= 1 &&
		p1.Icol <= p2.Icol+len(p2.Text) &&
		p2.Icol <= p1.Icol+len(p1.Text)
}

func getAdjacentNumbers(numbers, gears []Part) []Part {
	var result []Part
	for _, g := range gears {
		for _, n := range numbers {
			if isAdjacent(n, g) {
				result = append(result, n)
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
