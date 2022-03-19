package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	defer close(out)

	wg := new(sync.WaitGroup)

	for _, channel := range channels {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(channel)
	}

	wg.Wait()

	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-or(
		sig(5*time.Second),
		sig(2*time.Second),
		sig(9*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}

// done after 9.0075451s
