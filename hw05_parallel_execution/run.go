package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded   = errors.New("errors limit exceeded")
	ErrWorkersNumberInvalide = errors.New("workers number invalide")
	ErrTasksNumberInvalide   = errors.New("tasks number invalide")
	ErrNilTask               = errors.New("nil task")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrWorkersNumberInvalide
	}

	if len(tasks) == 0 {
		return ErrTasksNumberInvalide
	}

	if m < 0 {
		m = 0
	}

	var wg sync.WaitGroup
	tasksChan := make(chan Task)
	errChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())

	for n != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, tasksChan, errChan)
		}()
		n--
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(ctx, tasks, tasksChan)
	}()

	var tasksComplited int
	var errCount int
	var errorsLimitExceeded bool
	tasksNum := len(tasks)

	for taskResult := range errChan {
		if taskResult != nil {
			errCount++
			if errCount >= m {
				errorsLimitExceeded = true
				break
			}
		}
		tasksComplited++
		if tasksComplited == tasksNum {
			break
		}
	}

	cancel()
	wg.Wait()

	close(tasksChan)
	close(errChan)

	if errorsLimitExceeded {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func producer(ctx context.Context, tasks []Task, tasksChan chan<- Task) {
	for _, task := range tasks {
		select {
		case <-ctx.Done():
			return
		case tasksChan <- task:
		}
	}
}

func worker(ctx context.Context, tasksChan <-chan Task, errChan chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-tasksChan:
			doTask(ctx, task, errChan)
		}
	}
}

func doTask(ctx context.Context, task Task, errChan chan<- error) {
	var result error

	if task == nil {
		result = ErrNilTask
	} else {
		result = task()
	}

	select {
	case <-ctx.Done():
	case errChan <- result:
	}
}
