package worker

import (
	"gorchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

// TODO: implement CollectStats, RunTask, StartTask, StopTask functions
func (w *Worker) CollectStats() {

}

func (w *Worker) RunTask() {

}

func (w *Worker) StartTask() {

}

func (w *Worker) StopTask(t task.Task) {
	config := task.NewConfig(&t)
	d := task.NewDocker(config)

}
