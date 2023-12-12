package _legacy

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/DianaSun97/AdventGolang/common"
)

var input string

func main() {
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalf("read input file: %v", err)
	}
	input = fileContent

	sum := 0
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		line := parseNumbers(strings.TrimSpace(s))

		tower := generateTower(line)
		next := extrapolateNextValue(tower)

		sum += next
	}

	fmt.Println("Part 1 =", sum)
}

func extrapolateNextValue(tower [][]int) int {
	lastLine := len(tower) - 1
	tower[lastLine] = append(tower[lastLine], 0)
	for i := len(tower) - 2; i >= 0; i-- {
		idx := len(tower[i]) - 1
		tower[i] = append(tower[i], tower[i+1][idx]+tower[i][idx])
	}

	return tower[0][len(tower[0])-1]
}

func generateTower(num []int) [][]int {
	tower := [][]int{}
	tower = append(tower, num)
	for {
		row := tower[len(tower)-1]
		nextRow := []int{}
		finish := true

		for i := 0; i < len(row)-1; i++ {
			diff := row[i+1] - row[i]
			if diff != 0 {
				finish = false
			}
			nextRow = append(nextRow, diff)
		}

		tower = append(tower, nextRow)
		if finish {
			break
		}
	}

	return tower
}

func parseNumbers(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	numbers := []int{}
	for _, val := range re.FindAllStringSubmatch(s, -1) {
		num, _ := strconv.Atoi(val[0])
		numbers = append(numbers, num)
	}

	return numbers
}