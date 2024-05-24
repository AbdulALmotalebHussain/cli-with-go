// package main

// import (
// 	"fmt"
// )

// // Stack represents a stack data structure
// type Stack []string

// // NewStack creates a new empty stack
// func NewStack() Stack {
// 	return Stack{}
// }

// // IsEmpty checks if the stack is empty
// func (s Stack) IsEmpty() bool {
// 	return len(s) == 0
// }

// // Push adds an element to the top of the stack
// func (s *Stack) Push(item string) {
// 	*s = append(*s, item)
// 	fmt.Println("pushed item:", item)
// }

// func (s *Stack) Pop() (string, error) {

// 	if s.IsEmpty() {
// 		return "", fmt.Errorf("stack is empty")
// 	}
// 	item := (*s)[len(*s)-1]
// 	*s = (*s)[:len(*s)-1]
// 	fmt.Println("popped item:", item)
// 	return item, nil

// }

// func main() {

// 	Stack := NewStack()
// 	Stack.Push("1")
// 	Stack.Push("2")
// 	Stack.Push("kamak spanioal")

// 	poppedItem, err := Stack.Pop()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("popped items:",poppedItem)
// 		fmt.Println("stack items:",Stack)
// 	}
		
	
// }
