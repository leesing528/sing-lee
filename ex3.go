package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	leapYears := []int{}
	for i := x; i <= y; i++ {
		if isLeapYear(i) {
			leapYears = append(leapYears, i)
		}
	}

	leapYearsStr := make([]string, len(leapYears))
	for i, year := range leapYears {
		leapYearsStr[i] = strconv.Itoa(year)
	}

	fmt.Println(len(leapYears))
	fmt.Println(strings.Join(leapYearsStr, " "))
}
