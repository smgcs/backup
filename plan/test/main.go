package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	ch := make(chan int)
	go func() {
		ch <- 5
	}()
	go func() {
		fmt.Println(<-ch)
	}()
	time.Sleep(1e9)
}
