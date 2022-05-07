package goroutine_pool

import (
	"context"
	"log"
)

type TaskPool struct {
	canFunc  context.CancelFunc
	IsStop   bool
	Worker   chan func()
	Size     chan struct{}
	StopChan chan struct{}
}

func NewPool(ctx context.Context, size int) *TaskPool {
	return &TaskPool{
		Worker:   make(chan func()),
		Size:     make(chan struct{}, size),
		StopChan: make(chan struct{}),
	}
}

func (pool *TaskPool) AddTask(task func()) {
	if pool.IsStop {
		log.Println("TaskPool isStop!")
		return
	}

	select {
	case pool.Worker <- task:
		log.Println("add task")
	case pool.Size <- struct{}{}:
		go pool.worker()

	}
}

func (pool *TaskPool) Stop() {
	//pool.StopChan <- struct{}{}
	log.Println("stop!")
	pool.IsStop = true

	close(pool.Worker)
}
func (pool *TaskPool) ModerateStop() {
	log.Println("ModerateStop!")
	pool.IsStop = true
	pool.canFunc()

}
func (pool *TaskPool) worker() {
	defer func() {
		<-pool.Size
	}()

	for {
		select {
		case task, ok := <-pool.Worker:
			if !ok {
				return
			}
			if pool.IsStop {
				log.Println("isStop!!!")
				return
			}
			task()
		default:
			//	task()
			//	task = <-pool.Worker
		}
	}
}
