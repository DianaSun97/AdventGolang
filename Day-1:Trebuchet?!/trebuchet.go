package Day_1_Trebuchet__

import (
	"fmt"
	"github.com/DianaSun97/AdventGolang/common"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(strings.Trim(fileContent, " "), "\n")
	data := RemoveEmptyStrings(lines)

	totalChecksum, err := TotalChecksum(data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Total Checksum:", totalChecksum)
}

func calculateChecksum(chars []string) int {
	digits := []int{}

	for _, char := range chars {
		digit, err := strconv.Atoi(char)
		if err == nil {
			digits = append(digits, digit)
		}
	}

	if len(digits) == 0 {
		return 0
	}

	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	return firstDigit*10 + lastDigit
}

func TotalChecksum(input []string) (int, error) {
	startTime := time.Now()
	checksum := 0

	for _, line := range input {
		characters := strings.Split(line, "")
		checksum += calculateChecksum(characters)
	}

	log.Println("Checksum computation took:", time.Since(startTime))

	return checksum, nil
}

func RemoveEmptyStrings(lines []string) []string {
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return nonEmptyLines
}
