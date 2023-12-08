package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/DianaSun97/AdventGolang/common"
)

func main() {
	t := time.Now()
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalln(err)
	}

	part1, _ := PartOne(fileContent)
	fmt.Println("Part 1:", part1)

	part2, _ := PartTwo(strings.ReplaceAll(fileContent, " ", ""))
	fmt.Println("Part 2:", part2)

	fmt.Println(time.Since(t))
}

func PartOne(input string) (int64, error) {
	rows := strings.Split(input, "\n")
	times := Parse(rows[0])
	records := Parse(rows[1])

	res := int64(1)
	for i := 0; i < len(times); i++ {
		res *= WinningMoves(times[i], records[i])
	}
	return res, nil
}

func PartTwo(input string) (int64, error) {
	return PartOne(input)
}

func WinningMoves(time, record int64) int64 {
	// If we wait x ms, our boat moves `(time - x) * x` millimeters.
	// This breaks the record when `(time - x) * x > record`
	// or `-x^2  + time * x - record > 0`.

	// get the roots first
	x1, x2 := SolveEq(-1, time, -record)

	// integers in between the roots
	maxX := int64(math.Ceil(x2)) - 1
	minX := int64(math.Floor(x1)) + 1
	return maxX - minX + 1
}

func SolveEq(a, b, c int64) (float64, float64) {
	d := math.Sqrt(float64(b*b - 4*a*c))
	x1 := (-float64(b) - d) / (2 * float64(a))
	x2 := (-float64(b) + d) / (2 * float64(a))
	return math.Min(x1, x2), math.Max(x1, x2)
}

func Parse(input string) []int64 {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input, -1)
	var numbers []int64
	for _, match := range matches {
		num, _ := strconv.ParseInt(match, 10, 64)
		numbers = append(numbers, num)
	}
	return numbers
}
