package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file := "input.txt"
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	split := strings.Split(strings.TrimSpace(string(f)), "\n\n")

	cmp := func(a, b string) int {
		for _, s := range strings.Fields(split[0]) {
			if s := strings.Split(s, "|"); s[0] == a && s[1] == b {
				return -1
			}
		}
		return 0
	}

	run := func(sorted bool) (r int) {
		for _, s := range strings.Fields(split[1]) {
			if s := strings.Split(s, ","); slices.IsSortedFunc(s, cmp) == sorted {
				slices.SortFunc(s, cmp)
				n, _ := strconv.Atoi(s[len(s)/2])
				r += n
			}
		}
		return r
	}

	fmt.Println(run(true))
	fmt.Println(run(false))
}
