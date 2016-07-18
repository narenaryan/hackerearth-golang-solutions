package main

import "fmt"

// Implements Selection Sort in Go: O(n^2)
func SelectionSort(arr []int, length int) []int {
	var i, j, minIndex int
	i = 0
	for i < length-1 {
		minIndex = i
		j = i + 1
		for j < length {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			j += 1
		}
		if minIndex != i {
			temp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
		i += 1
	}
	return arr
}

// Implements Bubble Sort in Go: O(n^2)
func BubbleSort(arr []int, length int) []int {
	var i, j int
	i = 0
	for i < length-1 {
		j = i + 1
		for j < length {
			if arr[j] < arr[i] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
			j += 1
		}
		i += 1
	}
	return arr
}

// Implements Optimized Bubble Sort in Go: Best Case (O(n)))
func OptimizedBubbleSort(arr []int, length int) []int {
	var i, j int
	var swapped bool
	i = 0
	for i < length-1 {
		swapped = false
		j = i + 1
		for j < length {
			if arr[j] < arr[i] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
				swapped = true
			}
			j += 1
		}
		i += 1
		if !swapped {
			return arr
		}
	}
	return arr
}

// Implements Optimized Bubble Sort: O(n^2)
func InsertionSort(arr []int, length int) []int {
	var i, j int
	i = 1
	for i < length {
		j = i - 1
		insertElement := arr[i]
		for j >= 0 {
			if insertElement < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
			j -= 1
		}
		arr[j+1] = insertElement
		i += 1
	}
	return arr
}

func main() {
	b := make([]int, 5)
	b[0] = 12
	b[1] = 2
	b[2] = 32
	b[3] = 14
	b[4] = 7

	fmt.Println(BubbleSort(b, len(b)))
	fmt.Println(SelectionSort(b, len(b)))
	fmt.Println(InsertionSort(b, len(b)))
	fmt.Println(OptimizedBubbleSort(b, len(b)))

}
