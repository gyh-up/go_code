package main

import "fmt"

type Info interface {
	Getinfo() string
}
type Person struct {
	name string
	age int
}

func main() {
	p := Person{
		"guo",
		123,
	}
	printInfo(p)

}

func (p Person) Getinfo() string {
	return fmt.Sprintf("my name is %s,age is %d",p.name,p.age)
}

func printInfo(i Info) {
	fmt.Println(i.Getinfo())
}