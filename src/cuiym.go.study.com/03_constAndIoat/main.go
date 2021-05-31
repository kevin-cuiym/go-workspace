package main

import "fmt"

func main() {
	const pi = 3.1415
	const (
		a1 = iota
		a2
		a3
		_
		a4 = 100
		a5
		a6, a7 = iota + 1, iota + 2
		a8, a9 = iota + 1, iota + 2
	)

	fmt.Printf("a1:%d\n", a1)
	fmt.Printf("a2:%d\n", a2)
	fmt.Printf("a3:%d\n", a3)
	fmt.Printf("a4:%d\n", a4)
	fmt.Printf("a5:%d\n", a5)
	fmt.Printf("a6:%d\n", a6)
	fmt.Printf("a7:%d\n", a7)
	fmt.Printf("a8:%d\n", a8)
	fmt.Printf("a9:%d\n", a9)


}
