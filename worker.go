// Job worker.
package whitetree

import (
	"fmt"
	"sync"
)

type WorkerGroup struct {
	ErrChan chan error

	limiter  chan struct{}
	wg       sync.WaitGroup
	parsers  *ParserPackage
	handlers *HandlerPackage
}

func NewWorkerGroup(maxWorkers int, parsers *ParserPackage, handlers *HandlerPackage) *WorkerGroup {
	w := &WorkerGroup{
		ErrChan:  make(chan error),
		limiter:  make(chan struct{}, maxWorkers),
		parsers:  parsers,
		handlers: handlers,
	}

	return w
}

// Spawn a worker with task. Block if the number of running workers
// has reached maximum number.
func (w *WorkerGroup) Spawn(taskCtx *TaskContext) {
	w.limiter <- struct{}{}
	w.wg.Add(1)
	go func(taskCtx *TaskContext) {
		defer func() {
			if err := recover(); err != nil {
				w.ErrChan <- fmt.Errorf("execute failed: %q", err)
			}
			w.wg.Done()
			<-w.limiter
		}()
		if err := Dispatch(*taskCtx, *w.parsers, *w.handlers); err != nil {
			w.ErrChan <- err
		}
	}(taskCtx)
}

// Wait for all workers to complete.
func (w *WorkerGroup) Wait() error {
	w.wg.Wait()
	return <-w.ErrChan
}
