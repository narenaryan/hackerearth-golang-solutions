package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(bytestr []byte) string {
	i, j := 0, len(bytestr)-1
	for i <= j {
		first_char := string(bytestr[i])
		second_char := string(bytestr[j])
		if first_char != second_char {
			return "NO"
		}
		i += 1
		j -= 1
	}
	return "YES"
}

func main() {
	// Reading input
	inputreader := bufio.NewReader(os.Stdin)
	n, _ := inputreader.ReadString('\n')
	// Trim space
	str := strings.TrimSpace(n)
	bytestr := []byte(str)
	fmt.Println(isPalindrome(bytestr))
}
