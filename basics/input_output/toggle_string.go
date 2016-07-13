package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Reading input
	inputreader := bufio.NewReader(os.Stdin)
	n, _ := inputreader.ReadString('\n')
	// Trim space and then use Atoi
	str := strings.TrimSpace(n)
	bytestr := []byte(str)
	for index, character := range bytestr {
		char := string(character)
		if char == strings.ToUpper(char) {
			// Is lowercase
			bytestr[index] = []byte(strings.ToLower(char))[0]
		} else {
			bytestr[index] = []byte(strings.ToUpper(char))[0]
		}
	}
	fmt.Println(string(bytestr))
}
