package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	for {
		if value, ok := <-ch; ok {
			fmt.Println(value)
			if len(ch) == 0 {
				break
			}
		}
	}
}
