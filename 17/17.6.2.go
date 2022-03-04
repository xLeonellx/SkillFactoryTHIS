package main

import (
	"fmt"
	"time"
)

func main() {
	var ticker *time.Ticker = time.NewTicker(time.Second * 1)
	var t time.Time
	
	chanel1:=make(chan int)
	chanel2:=make(chan int)

	for {
		select {
		case message := <-chanel1:
			fmt.Println(message)
		case message := <-chanel2:
			fmt.Println(message)
		default:
			t = <-ticker.C
			outputMessage := []byte("Время: ")
			outputMessage = t.AppendFormat(outputMessage, "15:04:05")
			fmt.Println(string(outputMessage))
		}
	}
}
