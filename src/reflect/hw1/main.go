package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b any) {
	rVal := reflect.ValueOf(b)
	rType := reflect.TypeOf(b)
	fmt.Println("rVal: ", rVal)
	fmt.Println("rType: ", rType)
	fmt.Println("rVal.Kind(): ", rVal.Kind())
	fmt.Println("rType.Kind(): ", rType.Kind())
	iterf := rVal.Elem().Interface()
	fmt.Println("iterf: ", iterf)
	flo := iterf.(float64)
	fmt.Println("flo: ", flo)
}
func reflectTest2(b any) {

}

func main() {
	// var v float64 = 1.2
	// reflectTest(&v)
	var str string = "hello"
	fs := reflect.ValueOf(&str)
	fmt.Println("fs: ", fs.Elem())
	fs.Elem().SetString("jack")
	fmt.Println("fs: ", fs.Elem())
}
