package main

import "github.com/felipeas314/internal/stack"

func main() {
	s := stack.NewPersonStack()

	s.Push("AAAA,", "645.545.666.01")
	s.Push("BBBB,", "123.545.887.13")
	s.Push("CCCC,", "645.322.466.43")
	s.Push("DDDD,", "112.123.433.11")

	s.Display()

	s.Pop()

	s.Display()
}
