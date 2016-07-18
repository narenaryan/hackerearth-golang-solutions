package main

import "fmt"

// Recursive Fibonaacci O(2^n) Exponential runtime
func RecursiveFibonacci(num int) int {
	if num < 3 {
		return 1
	} else {
		return RecursiveFibonacci(num-1) + RecursiveFibonacci(num-2)
	}
}

// Iterative Fibonaacci O(n) Linear runtime
func IterativeFibonacci(num int) int {
	sum := 0
	first := 1
	second := 1
	count := 3
	for count <= num {
		sum = first + second
		first = second
		second = sum
		count += 1
	}
	return sum
}

func main() {
	fmt.Println(RecursiveFibonacci(15))
	fmt.Println(IterativeFibonacci(15))
}
