package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "input.txt"

	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	grid := map[image.Point]rune{}

	for y, line := range strings.Split(strings.TrimSpace(string(f)), "\n") {
		for x, letter := range line {
			grid[image.Point{x, y}] = letter
		}
	}

	part1 := 0
	part2 := 0

	part2String := "AMAMASASAMAMAS"

	for p := range grid {
		part1 += strings.Count(strings.Join(findWords(p, grid, 4), " "), "XMAS")
		part2 += strings.Count(part2String, strings.Join(findWords(p, grid, 2)[4:], ""))
	}

	fmt.Println(part1)
	fmt.Println(part2)

}

func findWords(p image.Point, grid map[image.Point]rune, l int) []string {

	directions := []image.Point{
		{0, -1},  // left
		{1, 0},   // down
		{0, 1},   // right
		{-1, 0},  // up
		{-1, -1}, // up-left
		{1, -1},  // down-left
		{1, 1},   // down-right
		{-1, 1},  // up-right
	}
	/*
		[]image.Point{
			{0, -1},  //left
			{0, 1},   //right
			{-1, 0},  //up
			{1, 0},   //down
			{-1, -1}, //up-left
			{1, -1},  //down-left
			{-1, 1},  //up-right
			{1, 1},   //down-right
		}
	*/

	words := make([]string, len(directions))

	for i, direction := range directions {
		for n := 0; n < l; n++ {
			words[i] += string(grid[p.Add(direction.Mul(n))])
		}
	}
	return words

}
