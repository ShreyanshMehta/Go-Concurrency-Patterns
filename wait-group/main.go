package main

import (
	"context"
	"k8s.io/apimachinery/pkg/util/rand"
	"sync"
	"time"
)

const (
	Tasks = 7
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < Tasks; i++ {
		wg.Add(1)
		task := NewTask(i, time.Duration(rand.Int63nRange(5000000000, 10000000000)))
		go func() {
			defer wg.Done()
			_ = task.Run(context.Background())
		}()
	}
	wg.Wait()
}
