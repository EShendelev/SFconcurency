package main

import (
	"fmt"
	"time"
)

func main() {
	first := make(chan int)
	second := make(chan int)

	for {
		select {
		case x := <-first:
			fmt.Printf("Message from channel №%d\n", x)
		case y := <-second:
			fmt.Printf("Message from channel №%d\n", y)
		default:
			time.Sleep(time.Second)
			fmt.Printf("\r%v", time.Now().Format("15:04:05"))
		}

	}
}
