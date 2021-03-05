package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Println("start")
	go func() {
		for i := 2; ; i++ {
			ch <- i
			fmt.Println("i:", i)
		}
	}()
	go func() {
		for {
			fmt.Println("<-ch:", <-ch)
		}
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("end")
}
