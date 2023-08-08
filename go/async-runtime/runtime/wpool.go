package runtime

import (
	"sync"

	"github.com/khekrn/async/domain/model"
)

type Executor struct {
	stateMachine *StateMachine
	channel      chan model.WorkflowRequest
	wg           sync.WaitGroup
	stopChannels []chan struct{}
	maxWorkers   int
}

func NewExecutor(maxWorkers int) *Executor {
	return &Executor{
		stateMachine: NewStateMachine(),
		channel:      make(chan model.WorkflowRequest),
		maxWorkers:   maxWorkers,
		stopChannels: make([]chan struct{}, maxWorkers),
	}
}

func (e *Executor) startWorker(workerID int) {
	e.wg.Add(1)
	stop := make(chan struct{})
	e.stopChannels[workerID] = stop

	go func() {
		defer e.wg.Done()
		for {
			select {
			case request, ok := <-e.channel:
				if !ok {
					return
				}
				e.stateMachine.Run(request)
			case <-stop:
				return
			}
		}
	}()
}

func (e *Executor) Start() {
	for i := 0; i < e.maxWorkers; i++ {
		e.startWorker(i)
	}
}

func (e *Executor) Stop() {
	for _, stop := range e.stopChannels {
		close(stop)
	}
	e.wg.Wait()
	close(e.channel)
}

func (e *Executor) ExecuteWorkflow(workflowRequest model.WorkflowRequest) {
	go func() {
		e.channel <- workflowRequest
	}()
}
