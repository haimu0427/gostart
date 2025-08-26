package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b any) {
	rtype := reflect.TypeOf(b)
	rvalue := reflect.ValueOf(b)
	fmt.Println("rtype:", rtype.Kind())
	fmt.Println("rvalue:", rvalue.Kind())
	n2 := rvalue.Int()
	fmt.Println(n2 + 100)
	iv := rvalue.Interface()
	n3, ok := iv.(int)
	if ok {
		fmt.Println("n3:", n3)
	}
	fmt.Println(n3 + 200)
	fmt.Println(iv)
}
func reflectTest2(b any) {

}

type Student struct {
	Name string
	Age  int
}

func main() {

	const PI = 3.14
	var num int = 100
	reflectTest(num)
	stu := Student{
		Name: "tom",
		Age:  18,
	}
	fmt.Println("reflectTest2:")
	reflectTest2(stu)
}
