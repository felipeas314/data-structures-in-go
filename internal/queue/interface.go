package queue

import "github.com/felipeas314/internal/domain"

type Queue interface {
	Enqueue(person domain.Person)
	Dequeue() *domain.Person
	Display()
	IsEmpty() bool
}
