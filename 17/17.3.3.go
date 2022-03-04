package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const step int64 = 5
const interationAmount int64 = 1000

func main() {
	var counter int64 = 0
	var c = sync.WaitGroup{}
	increment := func(i int) {
		defer c.Done()
		atomic.AddInt64(&counter, step)
	}
	var iterationCount int = int(interationAmount / step)
	for i := 1; i <= iterationCount; i++ {
		c.Add(1)
		go increment(i)
	}
	c.Wait()
	fmt.Println(counter)
}