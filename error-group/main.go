package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"
)

const (
	Tasks = 7
)

func main() {
	errGrp, errCtx := errgroup.WithContext(context.Background())
	for i := 0; i < Tasks; i++ {
		task := NewTask(i, time.Duration(rand.Int63nRange(5000000000, 10000000000)))
		errGrp.Go(func() error {
			return task.Run(errCtx)
		})
	}
	fmt.Printf("received err: %v", errGrp.Wait())
}
