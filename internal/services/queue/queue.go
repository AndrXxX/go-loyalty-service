package queue

import (
	"context"
	"errors"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"sync"
	"time"
)

type state struct {
	running  bool
	stopping bool
}

type runner struct {
	si      time.Duration
	workers []*worker
	s       state
	jobs    chan interfaces.QueueJob
	wg      sync.WaitGroup
}

func NewRunner(sleepInterval time.Duration) *runner {
	return &runner{
		si:      sleepInterval,
		workers: make([]*worker, 0),
		s:       state{running: false, stopping: false},
	}
}

func (r *runner) SetWorkersCount(count int) *runner {
	r.workers = make([]*worker, count)
	return r
}

func (r *runner) AddJob(j interfaces.QueueJob) error {
	if !r.s.running {
		return errors.New("trying to add a queue before starting runner")
	}
	r.jobs <- j
	return nil
}

func (r *runner) Run() error {
	if r.s.running {
		return errors.New("already running")
	}
	r.jobs = make(chan interfaces.QueueJob)
	logger.Log.Info("Queue running")
	r.s.running = true
	r.s.stopping = false
	go func() {
		for {
			for _, w := range r.workers {
				r.wg.Add(1)
				go func() {
					w.Process(r.jobs)
					r.wg.Done()
				}()
			}
			if r.s.stopping {
				r.wg.Done()
				r.s.stopping = false
				r.s.running = false
				logger.Log.Info("Queue stopped")
				return
			}
			r.wg.Wait()
			time.Sleep(r.si)
		}
	}()
	return nil
}

func (r *runner) Stop(ctx context.Context) error {
	select {
	default:
		logger.Log.Info("Queue shutting down")
		r.wg.Add(1)
		go func() {
			r.s.stopping = true
		}()
		r.wg.Wait()
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}
