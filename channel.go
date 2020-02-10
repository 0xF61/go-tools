package main

import (
	"fmt"
)

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	c := make(chan int)
	go strlen("Emir", c)
	go strlen("0xF61", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
