package service

import (
	"github.com/felipeas314/internal/domain"
	"github.com/felipeas314/internal/queue"
)

type QueueService struct {
	q queue.Queue
}

func NewQueueService(q queue.Queue) *QueueService {
	return &QueueService{q: q}
}

func (s *QueueService) AddPerson(name, email string) {
	s.q.Enqueue(domain.Person{Name: name, Email: email})
}

func (s *QueueService) RemovePerson() {
	s.q.Dequeue()
}

func (s *QueueService) ShowQueue() {
	s.q.Display()
}
