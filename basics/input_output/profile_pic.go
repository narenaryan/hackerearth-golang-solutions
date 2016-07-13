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
	cropSize, _ := strconv.Atoi(strings.TrimSpace(input))
	input, _ = inputreader.ReadString('\n')
	noOfCases, _ := strconv.Atoi(strings.TrimSpace(input))
	for i := 0; i < noOfCases; i++ {
		input, _ := inputreader.ReadString('\n')
		dimensions := cleanString(input, " ")
		xSize := dimensions[0]
		ySize := dimensions[1]

		if (xSize < cropSize) || (ySize < cropSize) {
			fmt.Println("UPLOAD ANOTHER")
		} else {
			if xSize == ySize {
				fmt.Println("ACCEPTED")
			} else {
				fmt.Println("CROP IT")
			}
		}
	}
}
