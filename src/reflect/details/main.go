package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b any) {
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal kind:", rVal.Kind())
	fmt.Println("rVal type:", rVal.Type())
	fmt.Println("rVal:", rVal)
	rVal.Elem().SetInt(20)
}
func main() {
	var num int = 10
	reflectTest(&num)
	fmt.Println("num=", num) // 20
}
