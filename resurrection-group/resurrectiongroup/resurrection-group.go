package resurrectiongroup

import "context"

type ResurrectionGroup struct {
	limiter chan struct{}
	cancel  context.CancelFunc
	ctx     context.Context
}

func WithContext(ctx context.Context, workerCount int) *ResurrectionGroup {
	ctx, cancel := context.WithCancel(ctx)
	return &ResurrectionGroup{
		limiter: make(chan struct{}, workerCount),
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (r *ResurrectionGroup) Run(f func(ctx context.Context)) {
	for {
		r.limiter <- struct{}{}
		if r.ctx.Err() != nil {
			return
		}
		go func() {
			f(r.ctx)
			<-r.limiter
		}()
	}
}

func (r *ResurrectionGroup) Stop() {
	r.cancel()
}
