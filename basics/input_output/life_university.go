package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read(inputreader *bufio.Reader) int {
	n, _ := inputreader.ReadString('\n')
	// Trim space and then use Atoi
	str := strings.TrimSpace(n)
	num, _ := strconv.Atoi(str)
	return num
}

func main() {
	// Read and Print until you found 42
	inputreader := bufio.NewReader(os.Stdin)
	fourtyTwo := read(inputreader)
	fmt.Println(fourtyTwo)
	for fourtyTwo != 42 {
		fourtyTwo = read(inputreader)
		if fourtyTwo != 42 {
			fmt.Println(fourtyTwo)
		}
	}
}
