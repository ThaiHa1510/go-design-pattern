package main

import (
	"fmt"
)

func Map[T any, U any](input []T, mapper func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

func CoverStr[T any](input []T, cover func(T) T) []T {
	result := make([]T, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = cover(input[i])
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	output := Map(input, func(i int) int64 {
		return int64(i * 2)
	})
	fmt.Println(output)
	fmt.Println("--- Stack ---")
	var stack Stack[int]
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println("--- CoverStr ---")
	inputStr := []string{"a", "b", "c", "d", "e"}
	outputStr := CoverStr(inputStr, func(s string) string {
		return s + "!"
	})
	fmt.Println(inputStr)
	fmt.Println(outputStr)
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(newEle T) {
	s.elements = append(s.elements, newEle)
}

func (s *Stack[T]) Pop() T {
	last := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return last
}
