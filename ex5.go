package main

import "fmt"

func main() {
	numbers := make([]int, 50)
	for i := range numbers {
		numbers[i] = i + 1
	}

	var result []int
	for _, num := range numbers {
		if num%3 != 0 {
			result = append(result, num)
		}
	}
	result = append(result, 114514)
	fmt.Println(result)
}
