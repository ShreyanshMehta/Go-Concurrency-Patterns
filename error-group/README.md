# Error Group
This is an implementation of GO's [`errgroup` package](https://pkg.go.dev/golang.org/x/sync/errgroup).
Package `errgroup` provides synchronization, error propagation, and Context cancelation for groups of goroutines working on subtasks of a common task.

---
```text
type Group
    func WithContext(ctx context.Context) (*Group, context.Context)
    func (g *Group) Go(f func() error)
    func (g *Group) SetLimit(n int)
    func (g *Group) Wait() error
```
 
## Problem it solves

Let's say we have multiple subtasks working inside a goroutine to fulfill a common task. Now, one task got terminated because of some error. 
Sometimes, it doesn't make any sense to run other sibling routines. In such scenarios, we are expected to terminate the whole application by closing other goroutines as well and log the error which caused this termination.

In this type of scenario, we can use `errgroups` to handle routines. `errgroup` can also be used to limit the goroutines running concurrently.  
