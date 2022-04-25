package goroutine_pool

import (
	"context"
	"log"
)

type TaskPool struct {
	Ctx      context.Context
	canFunc  context.CancelFunc
	IsStop   bool
	Worker   chan func()
	Size     chan struct{}
	StopChan chan struct{}
}

func NewPool(ctx context.Context, size int) *TaskPool {
	poolCtx, cancel := context.WithCancel(ctx)
	return &TaskPool{
		Ctx:      poolCtx,
		canFunc:  cancel,
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
	ctx, _ := context.WithCancel(pool.Ctx)

	select {
	case pool.Worker <- task:
		log.Println("add task")
	case pool.Size <- struct{}{}:
		go pool.worker(ctx)

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
func (pool *TaskPool) worker(ctx context.Context) {
	defer func() {
		<-pool.Size
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Doneee!")
			return
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
