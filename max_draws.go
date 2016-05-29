package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputreader *bufio.Reader

func main() {
	inputreader = bufio.NewReader(os.Stdin)
	n, _ := inputreader.ReadString('\n')
	// Trim space and then use Atoi
	no_of_iters, _ := strconv.Atoi(strings.TrimSpace(n))
	for i := 0; i < no_of_iters; i++ {
		stream, _ := inputreader.ReadString('\n')
		trimmed_stream := strings.TrimSpace(stream)
		split_arr := strings.Split(trimmed_stream, " ")

    fmt.Println((2 * qx - px), (2 * qy - py))
	}
}
