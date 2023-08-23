# Error Group
This is an implementation of GO's [`waitgroup` package](https://pkg.go.dev/golang.org/x/sync/errgroup).
A `WaitGroup` waits for a collection of goroutines to finish.

---
```text
type WaitGroup
    Add(delta int)
    Done()
    Wait()
```
 
## Problem it solves

Let's say we have multiple goroutine running. Now, you want to block a goroutine until the collection of goroutine finish their job.

The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
