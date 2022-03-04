package main

import (
	"fmt"
	"time"
)

func main() {

	chanel1:=make(chan int)
	chanel2:=make(chan int)
	go func() {
		for  {
		time.Sleep(time.Second)
		chanel1 <- 1
		}
	}()
	go func() {
		for  {
		time.Sleep(time.Second * 2)
		chanel2 <- 2
		}
	}()
	for {
		select {
		case message := <-chanel1:
			fmt.Println(message)
		case message := <-chanel2:
			fmt.Println(message)
		}
	}
}
