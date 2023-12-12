package _legacy

import (
	"fmt"
	"github.com/DianaSun97/AdventGolang/common"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	WinningNumbers []int
	YourNumbers    []int
}

func main() {
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(fileContent, "\n")
	totalPoints := 0
	for i, line := range lines {
		if i+1 < len(lines) {
			card := parseCard(line)
			points := calculatePoints(card)
			totalPoints += points
			fmt.Printf("Card %d: %d points\n", i+1, points)
		}
	}

	fmt.Printf("Total Points: %d\n", totalPoints)
}

func parseCard(line string) Card {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		log.Fatalf("Invalid input line: %s", line)
	}
	return Card{
		WinningNumbers: extractNumbers(parts[0]),
		YourNumbers:    extractNumbers(parts[1]),
	}
}

func extractNumbers(s string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	numbers := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		numbers[i] = num
	}
	return numbers
}

func calculatePoints(card Card) int {
	points := 0
	for _, winningNumber := range card.WinningNumbers {
		if contains(card.YourNumbers, winningNumber) {
			points++
		}
	}

	return int(math.Pow(2, float64(points)))
}

func contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}
