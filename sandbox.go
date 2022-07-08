package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 64
	close(c)
	// i, r := <-c
	// d, e := <-c
	// fmt.Println(i, r)
	// fmt.Println(d, e)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
