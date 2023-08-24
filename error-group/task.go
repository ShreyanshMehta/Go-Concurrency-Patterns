package main

import (
	"context"
	"fmt"
	"time"
)

type Task struct {
	taskID    int
	errPeriod time.Duration
}

func NewTask(id int, errPeriod time.Duration) *Task {
	return &Task{
		taskID:    id,
		errPeriod: errPeriod,
	}
}

func (t *Task) Run(ctx context.Context) error {
	ticker := time.NewTimer(t.errPeriod)
	fmt.Printf("running task %v\n", t.taskID)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stopping task %v\n", t.taskID)
			return ctx.Err()
		case <-ticker.C:
			fmt.Printf("stopping task %v\n", t.taskID)
			return fmt.Errorf("some error from task %v", t.taskID)
		}
	}
}
