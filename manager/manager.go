package manager

import (
	"gorchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue // tasks which have been submitted by user but have yet to be assigned to a worker
	TaskDb        map[string][]task.Task
	EventDb       map[string][]task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

// TODO: implement SelectWorker, UpdateTasks, SendWork functions
func (m *Manager) SelectWorker() {

}

func (m *Manager) UpdateTasks() {

}

func (m *Manager) SendWork() {

}
