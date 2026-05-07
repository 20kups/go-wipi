package models

// ------ Task Type ------
type TaskType string

const (
	TaskTypeProcess TaskType = "process"
	TaskTypeLatchAndClear TaskType = "latch-and-clear"
)
// ------ /Task Type ------

// ------ Task ------
type Task struct {
	taskType TaskType
	fn   func()
}

func NewTask(tt TaskType, fn func()) *Task {
	return &Task{taskType: tt, fn: fn}
}

func (t *Task) Type() TaskType {
	return t.taskType
}

func (t *Task) Run() {
	t.fn()
}
// ------ /Task ------

// ------ Task Queue ------
type TaskQueue struct {
	ch chan *Task
}

func NewTaskQueue(capacity int) *TaskQueue {
	return &TaskQueue{ch: make(chan *Task, capacity)}
}

func (tq *TaskQueue) Enqueue(t *Task) {
	tq.ch <- t
}

func (tq *TaskQueue) Dequeue() (*Task, bool) {
	t, ok := <-tq.ch
	return t, ok
}
// ------ /Task Queue ------
