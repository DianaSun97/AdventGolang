package Day_1_Trebuchet__

import (
	"bufio"
	"fmt"
	"github.com/DianaSun97/AdventGolang/common"
	"log"
	"strconv"
	"strings"
	"time"
)

func Solve(input string) (string, string) {

	part2 := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		{
			lineAsBytes := []byte(line)

			replace := func(lineAsBytes []byte, oldValue string, newValue byte) {
				for {
					index := strings.Index(string(lineAsBytes), oldValue)
					if index < 0 {
						return
					}
					lineAsBytes[index+(len(oldValue)/2)] = newValue
					lineAsBytes = lineAsBytes[index+1:]
				}
			}

			replace(lineAsBytes, "one", '1')
			replace(lineAsBytes, "two", '2')
			replace(lineAsBytes, "three", '3')
			replace(lineAsBytes, "four", '4')
			replace(lineAsBytes, "five", '5')
			replace(lineAsBytes, "six", '6')
			replace(lineAsBytes, "seven", '7')
			replace(lineAsBytes, "eight", '8')
			replace(lineAsBytes, "nine", '9')

			tens := int(lineAsBytes[strings.IndexAny(string(lineAsBytes), "0123456789")]-'0') * 10
			ones := int(lineAsBytes[strings.LastIndexAny(string(lineAsBytes), "0123456789")] - '0')
			part2 += tens + ones
		}
	}

	return strconv.Itoa(part2), ""
}

func main() {
	t := time.Now()
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalln(err)
	}

	part2, _ := Solve(fileContent)
	fmt.Println("Part 2:", part2, time.Since(t))
}
