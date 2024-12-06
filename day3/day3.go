package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"

	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	value, count := multiply(string(f))

	fmt.Println(count)
	fmt.Println(value)

	fmt.Println(exclude(string(f)))

}

func multiply(s string) (r int, count int) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, m := range re.FindAllStringSubmatch(s, -1) {
		n1, _ := strconv.Atoi(m[1])
		n2, _ := strconv.Atoi(m[2])
		r += n1 * n2
		count++
	}
	return r, count
}

func exclude(s string) (r int) {
	re := regexp.MustCompile(`(?s)(?:^|do\(\)).*?(?:don't\(\)|$)`)
	r, _ = multiply(strings.Join(re.FindAllString(string(s), -1), ""))
	return r
}
