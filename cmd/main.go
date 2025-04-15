package main

import (
	"github.com/felipeas314/internal/queue"
	"github.com/felipeas314/internal/service"
)

func main() {
	q := queue.NewLinkedListQueue()
	qService := service.NewQueueService(q)

	qService.AddPerson("John Doe", "john@example.com")
	qService.AddPerson("Jane Smith", "jane@example.com")
	qService.AddPerson("Carlos Silva", "carlos@example.com")

	qService.ShowQueue()

	qService.RemovePerson()
	qService.ShowQueue()
}
