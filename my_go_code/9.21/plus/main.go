package main
import(
	"fmt"
)

func main(){
	n := PlusTwo()(99)
	fmt.Printf("99+2=%v\n", n)
	num := PlusX(99)(100)
	fmt.Printf("99+100=%v\n", num)
}

func PlusTwo() func(int) int {
	return func(x int) int {return x+2}
}

func PlusX(x int) func(int) int{
	return func(y int) int {return x+y}
}