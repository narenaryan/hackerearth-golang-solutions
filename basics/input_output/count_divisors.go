package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cleanString(stream string, seperator string) []int {
	// Trims the stream and then splits
	trimmed_stream := strings.TrimSpace(stream)
	split_arr := strings.Split(trimmed_stream, seperator)
	// convert strings to integers and store them in a slice
	clean_arr := make([]int, len(split_arr))

	for index, val := range split_arr {
		clean_arr[index], _ = strconv.Atoi(val)
	}
	return clean_arr
}

func main() {
	// Reading input
	inputreader := bufio.NewReader(os.Stdin)
	input, _ := inputreader.ReadString('\n')
	// Trim space and then use Atoi
	inputs := cleanString(input, " ")
	l, r, k, count := inputs[0], inputs[1], inputs[2], 0
	for i := l; i <= r; i++ {
		if i%k == 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
