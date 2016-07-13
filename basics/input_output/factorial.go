package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	// Reading input
	inputreader := bufio.NewReader(os.Stdin)
	input, _ := inputreader.ReadString('\n')
	// Trim space and then use Atoi
	num, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Println(factorial(num))
}
