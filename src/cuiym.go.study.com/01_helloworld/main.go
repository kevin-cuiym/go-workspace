package main

import "fmt"

var userName string

func main() {
	fmt.Println("Hello world!!")

	var (
		id string
		age int
		isOk bool
	)

	id = "Cui yong min"
	age = 10
	isOk = true

	fmt.Printf("name:%s\n", id)
	fmt.Println(age)
	fmt.Println(isOk)

}