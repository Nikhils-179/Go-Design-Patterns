//Fan-in Fanout concurrency design pattern code

// Fan-in
package main

import (
	"fmt"
	"sync"
)

func Fanin[T any](channels []<-chan T) <-chan T {
	out := make(chan T)

	wg := &sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan T) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func Fanout[T any](source chan T, n int) []<-chan T {

	out := make([]<-chan T, 0, n)

	for i := 0; i < n; i++ {
		ch := make(chan T)
		out = append(out, ch)

		go func() {
			defer close(ch)
			for v := range source {
				ch <- v
			}
		}()
	}
	return out
}

func main() {
	source := make(chan any, 10)
	for i := 0; i < 10; i++ {
		source <- i
	}

	close(source)
	outputchannels := Fanout(source, 4)
	result := Fanin(outputchannels)
	for v := range result {
		fmt.Println(v)
	}

}
