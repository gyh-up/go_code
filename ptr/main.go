package main

import "fmt"

func main()  {
	arr := [3]int{1,2,3}
	fmt.Println(&arr[0],&arr[1],&arr[2])
	fmt.Println(arr)
}
