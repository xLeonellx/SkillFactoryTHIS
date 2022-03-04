package main

import (
	"fmt"
	"sync"
)

const maxCount = 100
func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	chanel:=make(chan int)

	go func() {
		defer wg.Done()
		for i:=1;i<=maxCount;i++{
			chanel<-i
		}
		close(chanel)
	}()
	go func() {
		defer wg.Done()
		for i:=range chanel{
			fmt.Printf("%v ", i)
		}
	}()
	wg.Wait()
}
