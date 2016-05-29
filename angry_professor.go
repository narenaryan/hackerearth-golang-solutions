package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputreader *bufio.Reader

func cleanString(stream string, seperator string) []int{
  // Trims the stream and then splits
  trimmed_stream := strings.TrimSpace(stream)
  split_arr := strings.Split(trimmed_stream, seperator)
  // convert strings to integers and store them in a slice
  clean_arr := make([]int, len(split_arr))
  for index,val := range split_arr{
    clean_arr[index], _ = strconv.Atoi(val)
  }
  return clean_arr
}


func main() {
  // Reading input
	inputreader = bufio.NewReader(os.Stdin)
	n, _ := inputreader.ReadString('\n')
	// Trim space and then use Atoi
	no_of_iters, _ := strconv.Atoi(strings.TrimSpace(n))
	for i := 0; i < no_of_iters; i++ {
		stream, _ := inputreader.ReadString('\n')
    students_arr := cleanString(stream, " ")
    cancellation_point := students_arr[1]
    // convert strings to integers and store them in a slice
    stream, _  = inputreader.ReadString('\n')
		attendence_arr := cleanString(stream, " ")
    count := 0
    for _,val := range attendence_arr{
      if val <= 0{
        count++
      }
    }
    if count >= cancellation_point {
      fmt.Println("NO")
    } else{
      fmt.Println("YES")
    }
	}
}
