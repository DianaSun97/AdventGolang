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

type Game struct {
	rounds []Set
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

func (s *Set) isValid(validBag *Set) bool {
	if s.red > validBag.red {
		return false
	} else if s.blue > validBag.blue {
		return false
	} else if s.green > validBag.green {
		return false
	}
	return true
}

func parseGame(gameLine string) *Game {
	sets := strings.Split(gameLine, ": ")
	setsSplit := strings.Split(sets[1], ";")

	rounds := make([]Set, len(setsSplit))
	game := Game{rounds: rounds}

	for idx, setStr := range setsSplit {
		set := parseSet(setStr)
		game.rounds[idx] = *set
	}

	return &game
}

func (g *Game) isValid(validBag *Set) bool {
	for _, round := range g.rounds {
		if !round.isValid(validBag) {
			return false
		}
	}
	return true
}

func main() {
	fileContent, err := common.ReadInputFile()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(fileContent))
	validBag := Set{red: 12, green: 13, blue: 14}
	idSum := 0
	gameNr := 1

	for scanner.Scan() {
		gameLine := scanner.Text()
		game := parseGame(gameLine)

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
