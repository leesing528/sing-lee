package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	heights := strings.Split(scanner.Text(), " ")
	appleHeights := make([]int, 10)
	for i, heightStr := range heights {
		appleHeights[i], _ = strconv.Atoi(heightStr)
	}
	scanner.Scan()
	taoTaoHeight, _ := strconv.Atoi(scanner.Text())
	taoTaoHeight += 30

	count := 0
	for _, height := range appleHeights {
		if height <= taoTaoHeight {
			count++
		}
	}
	fmt.Println(count)
}
