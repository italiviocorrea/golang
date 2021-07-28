package worker

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/italiviocorrea/golang/mycorch.io/cube/cmd/task"
)
type Worker struct {
	Name string
	Queue queue.Queue
	Db map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats()  {
	fmt.Println("Estou coletando status")
}

func (w *Worker) RunTask()  {
	fmt.Println("Estou iniciando ou parando a tarefa")
}

func (w *Worker) StartTask()  {
	fmt.Println("Estou iniciando a tarefa")
}

func (w *Worker) StopTask() {
	fmt.Println("Estou parando a tarefa")
}

