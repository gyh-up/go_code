package main

import (
	"fmt"
	"reflect"
)

func main () {
	var flo float64
	flo = 1.2
	reflectFloat(flo)
}


func reflectFloat(i interface{}){
	refType := reflect.TypeOf(i)
	refVal := reflect.ValueOf(i)

	fmt.Println(refType, "-", refVal)
	fmt.Printf("%T\n", refType)

	value := refVal.Float()

	fmt.Printf("%f-%T", value, value)

}

