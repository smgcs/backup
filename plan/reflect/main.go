package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//quote_change()
	dynamic_call()
}

type User struct {
	Username string
	age      int
}

func quote_change() {
	u := &User{"Jack", 23}
	p := reflect.ValueOf(u).Elem()
	p.FieldByName("Username").SetString("Tom")
	f := p.FieldByName("age")
	fmt.Println(f.CanSet())
	// 判断是否能获取地址。
	if f.CanAddr() {
		age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		// age := (*int)(unsafe.Pointer(f.Addr().Pointer())) // 等同
		*age = 88
	}
	// 注意 p 是 Value 类型，需要还原成接⼝口才能转型。
	fmt.Println(u, p.Interface().(User))
}

type Data struct {
	one string
}

func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}
	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name)
	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf("  in[%d] %v\n", i, t.In(i))
	}
	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf("  out[%d] %v\n", i, t.Out(i))
	}
}
func dynamic_call() {
	d := new(Data)
	v := reflect.ValueOf(d)
	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		out := m.Call(in)
		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}
	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
	fmt.Println("-----------------------")
	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
}
