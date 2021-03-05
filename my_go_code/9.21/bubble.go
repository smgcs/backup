package main

import "fmt"

func main() {
	array := []int{1, 4, 7, 9, 7, 6, 4, 3}
	fmt.Printf("unsored array: %v\n", array)
	Bubble(array)
	fmt.Printf("sored array: %v\n", array)

}

func Bubble(l []int) {
	for i := 0; i < len(l); i++ {
		for j := 0; j < i; j++ {
			if l[j] < l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}
