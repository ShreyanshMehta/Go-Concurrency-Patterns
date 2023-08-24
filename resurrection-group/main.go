package main

import (
	"context"
	"fmt"
	"github.com/ShreyanshMehta/go-concurrency-patterns/resurrection-group/resurrectiongroup"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/uuid"
	"time"
)

const (
	workers = 3
)

func main() {
	resurrectGrp := resurrectiongroup.WithContext(context.Background(), workers)

	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("stopping the resurrection group")
		resurrectGrp.Stop()
	}()

	resurrectGrp.Run(func(ctx context.Context) {
		task := NewTask(string(uuid.NewUUID()), time.Duration(rand.Int63nRange(5000000000, 10000000000)))
		_ = task.Run(ctx)
	})
}
