package manager

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/italiviocorrea/golang/mycorch.io/cube/cmd/task"
)

type Manager struct {
	Pending queue.Queue
	TaskDb map[string][]task.Task
	EventDb map[string][]task.TaskEvent
	Workers []string
	WorkerTaskMap map[string]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("Estou selecionando um trabalhador adequado")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("Estou atualizando a tarefa")
}

func (m *Manager) SendWork() {
	fmt.Println("Estou enviando um trabalho para o trabalhador")
}

