package main

import (
	"context"
	"sync"
)

type Group struct {
	ctx          context.Context
	cancel       context.CancelCauseFunc
	wg           sync.WaitGroup
	limitHandler chan struct{}
}

func WithContext(ctx context.Context) (*Group, context.Context) {
	errCtx, cancel := context.WithCancelCause(ctx)
	return &Group{
		ctx:    errCtx,
		cancel: cancel,
		wg:     sync.WaitGroup{},
	}, errCtx
}

func (g *Group) SetLimit(n int) {
	g.limitHandler = make(chan struct{}, n)
}

func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		g.limitHandler <- struct{}{}
		defer g.wg.Done()
		if err := f(); err != nil {
			g.cancel(err)
		}
		<-g.limitHandler
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	return context.Cause(g.ctx)
}
