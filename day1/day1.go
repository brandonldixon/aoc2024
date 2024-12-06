package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {

	//filename
	filename := "lists.csv"

	records := parseLists(filename)

	var left, right []int

	for _, record := range records {
		val1, err1 := strconv.Atoi(record[0])
		val2, err2 := strconv.Atoi(record[1])
		if err1 != nil || err2 != nil {
			log.Fatal(err1, err2)
		}

		left = append(left, val1)
		right = append(right, val2)
	}

	slices.Sort(left)
	slices.Sort(right)

	compare := compareLists(left, right)

	var total int

	for _, v := range compare {
		total += v
	}

	fmt.Println(total)

	similarityScore := similarityScore(left, right)

	fmt.Println(similarityScore)

}

func compareLists(l, r []int) []int {
	diff := make([]int, len(l))

	if len(l) == len(r) {
		for i := 0; i < len(l); i++ {
			if l[i] >= r[i] {
				diff[i] = l[i] - r[i]
			}
			if r[i] > l[i] {
				diff[i] = r[i] - l[i]
			}
		}
	}
	return diff
}

func parseLists(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	lists, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return lists
}

func similarityScore(l, r []int) int {
	totalCounter := 0

	for _, leftNum := range l {
		counter := 0
		for _, rightNum := range r {
			if leftNum == rightNum {
				counter += 1
			}
		}
		totalCounter += leftNum * counter
	}
	return totalCounter
}
