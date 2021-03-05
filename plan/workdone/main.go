package main

import "fmt"

func main() {
	done := make(chan bool, 1)
	go work(done)
	<-done
	fmt.Println("waiting for you work done")
}

func work(done chan bool) {
	//defer func(){
	//done <- true
	//}
	fmt.Println("work done")
	done <- true
}
