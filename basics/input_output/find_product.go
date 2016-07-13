package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func cleanString(stream string, seperator string, num int) []int {
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
	num, _ := strconv.Atoi(strings.TrimSpace(input))
	arr, _ := inputreader.ReadString('\n')
	intArr := cleanString(arr, " ", num)

	answer := 1

	for _, item := range intArr {
		answer = (answer * item) % (int(math.Pow(float64(10), 9)) + 7)
	}
	fmt.Println(answer)
}
