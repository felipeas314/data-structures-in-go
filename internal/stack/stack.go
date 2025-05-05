package stack

import "fmt"

type PersonStack struct {
	top *PersonNode
}

func NewPersonStack() *PersonStack {
	return &PersonStack{}
}

func (s *PersonStack) Push(name, cpf string) {
	newNode := &PersonNode{
		Name: name,
		CPF:  cpf,
		Next: s.top,
	}

	s.top = newNode
	fmt.Printf("%s has been pushed onto the stack.\n", name)
}

func (s *PersonStack) Pop() *PersonNode {
	if s.top == nil {
		fmt.Println("The stack is empty")
		return nil
	}

	removed := s.top
	s.top = s.top.Next

	fmt.Printf("%s has been popped from the stack.\n", removed.Name)

	return removed
}

func (s *PersonStack) Display() {
	if s.top == nil {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Print("People in the stack (Top to bottom):")
	current := s.top
	for current != nil {
		fmt.Printf("- Name: %s, CPF: %s\n", current.Name, current.CPF)
		current = current.Next
	}
}
