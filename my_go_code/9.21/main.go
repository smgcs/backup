package main

import "fmt"
import "strconv"

type stack struct{
	i int
	data [10]int
}

func (s *stack)push(v int){
	if s.i > 9{
		return
	}
	s.data[s.i] = v
	s.i++
	// fmt.Println(*s)
}

func (s *stack)pop(){
	if s.i==0{
		fmt.Println("can't pop, it's null")
		return
	}
	s.i = s.i-1
	s.data[s.i] = 0
}

func (s stack)String() string{
	var str = ""
	for i:=0;i<s.i;i++{
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]" 
	}
	return str
}
/*
func main(){
	var s1 stack
	for i := 0; i<=5; i++{
		s1.push(i)
	}
	fmt.Println(s1.String())
	s1.pop()
	fmt.Println(s1)
}*/