package main

import "math"
import "fmt"

// Linear Search algorithm O(n)
func LinearSearch(arr []int, length int, elem int) int {
	for i := 0; i < length; i++ {
		if arr[i] == elem {
			return i
		}
	}
	return -1
}

// Binary Search Algorithm O(log n)
func BinarySearch(arr []int, length int, elem int) int {
	left := 0
	right := length - 1

	for left <= right {
		mid := int(math.Floor((float64(left) + float64(right)) / 2))
		if elem == arr[mid] {
			return mid
		} else if elem < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	b := make([]int, 5)
	b[0] = 2
	b[1] = 7
	b[2] = 14
	b[3] = 32
	b[4] = 41
	fmt.Println(LinearSearch(b, len(b), 2))
	fmt.Println(BinarySearch(b, len(b), 2))
	fmt.Println(LinearSearch(b, len(b), 14))
	fmt.Println(BinarySearch(b, len(b), 14))
	fmt.Println(LinearSearch(b, len(b), 41))
	fmt.Println(BinarySearch(b, len(b), 41))
}
