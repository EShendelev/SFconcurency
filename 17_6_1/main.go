package main

import (
	"fmt"
	"time"
)

func main() {
	first := make(chan int)
	second := make(chan int)

go func() {
	for {
		time.Sleep(500*time.Millisecond)
		first <- 1
	}
}()

go func() {
	for {
		time.Sleep(1000*time.Millisecond)
		second <- 2
	}
}()

for {
	select {
	case x := <-first:
		fmt.Printf("Message from channel №%d\n", x)
	case y := <-second:
		fmt.Printf("Message from channel №%d\n", y)
	}
}
}