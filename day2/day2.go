package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	var codes [][]int
	var part1 int
	var part2 int

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		var intSlice []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			intSlice = append(intSlice, num)
		}
		codes = append(codes, intSlice)
	}

	for _, list := range codes {
		if check(list) {
			part1++
		}
	}
	fmt.Println(part1)

	for _, list := range codes {
		for i := range list {
			if check(slices.Delete(slices.Clone(list), i, i+1)) {
				part2++
				break
			}
		}
	}

	fmt.Println(part2)
}

func check(r []int) bool {
	for i := 1; i < len(r); i++ {
		if d := r[i] - r[i-1]; d*(r[1]-r[0]) <= 0 || d < -3 || d > 3 {
			return false
		}
	}
	return true
}
