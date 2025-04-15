package queue

import (
	"fmt"

	"github.com/felipeas314/internal/domain"
)

type node struct {
	person domain.Person
	next   *node
}

type LinkedListQueue struct {
	head *node
	tail *node
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (q *LinkedListQueue) Enqueue(person domain.Person) {
	newNode := &node{person: person}

	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

	fmt.Printf("%s has entered the queue.\n", person.Name)
}

func (q *LinkedListQueue) Dequeue() *domain.Person {
	if q.head == nil {
		fmt.Println("The queue is empty!")
		return nil
	}

	removed := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	fmt.Printf("%s has been removed from the queue.\n", removed.person.Name)
	return &removed.person
}

func (q *LinkedListQueue) Display() {
	if q.head == nil {
		fmt.Println("The queue is empty!")
		return
	}

	fmt.Println("People in the queue:")
	current := q.head
	for current != nil {
		fmt.Printf("Name: %s, Email: %s\n", current.person.Name, current.person.Email)
		current = current.next
	}
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.head == nil
}
