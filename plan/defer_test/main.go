package main

import (
	"fmt"
)

func f() (ret int) {
	//defer 在return之后执行
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(f()) //print 2
}
