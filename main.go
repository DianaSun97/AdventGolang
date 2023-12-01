package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	//"github.com/your-custom-path/common"
)

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

func ComputeTotalChecksum(input []string) (int, error) {
	startTime := time.Now()
	checksum := 0

	for _, line := range input {
		characters := strings.Split(line, "")
		checksum += calculateChecksum(characters)
	}

	log.Println("Checksum computation took:", time.Since(startTime))

	return checksum, nil
}

func main() {
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(strings.Trim(fileContent, " "), "\n")
	data := common.RemoveEmptyStrings(lines)

	totalChecksum, err := ComputeTotalChecksum(data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Total Checksum:", totalChecksum)
}
