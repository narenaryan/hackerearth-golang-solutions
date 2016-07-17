package main

import "fmt"

type Stack struct {
	Top int
	Arr []int
}

func (s Stack) init(size int) {
	s.Top = 1
	s.Arr = make([]int, size)
	return
}

func (s Stack) push(item int) {
	s.Arr = append(s.Arr, item)
	s.Top += 1
}

func (s Stack) pop() int {
	if s.Top == -1 {
		fmt.Println("Stack is empty")
		return 0
	}
	arrLength := len(s.Arr) - 1
	poppedElement := s.Arr[arrLength]
	s.Arr = s.Arr[:arrLength]
	return poppedElement
}

func main() {
	myStack := Stack()
	myStack.init(5)
	myStack.push(12)
	fmt.Println(myStack.pop())
}
