package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DianaSun97/AdventGolang/common"
)

type Set struct {
	red   int
	green int
	blue  int
}

func parseSet(setStr string) *Set {
	splits := strings.Split(setStr, ",")

	red := 0
	green := 0
	blue := 0

	for _, split := range splits {
		splitDraw := strings.Split(strings.TrimSpace(split), " ")
		number, err := strconv.Atoi(splitDraw[0])
		if err != nil {
			log.Fatalln("Unable to parse set", split)
		}
		color := splitDraw[1]

		if color == "red" {
			red = number
		} else if color == "green" {
			green = number
		} else if color == "blue" {
			blue = number
		}
	}

	return &Set{red: red, green: green, blue: blue}
}

func (s *Set) isValid(bag *Set) bool {
	if s.red > bag.red {
		return false
	} else if s.blue > bag.blue {
		return false
	} else if s.green > bag.green {
		return false
	}
	return true
}

func main() {
	fileContent, err := common.ReadInputFile()

	if err != nil {
		log.Fatalln(err)
	}

	reader := strings.NewReader(fileContent)
	scanner := bufio.NewScanner(reader)
	validBag := Set{red: 12, green: 13, blue: 14}
	idSum := 0
	gameNr := 1

	for scanner.Scan() {
		gameLine := scanner.Text()
		game := parseSet(gameLine)

		if game.isValid(&validBag) {
			idSum += gameNr
		}

		gameNr++
	}

	fmt.Println(idSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
