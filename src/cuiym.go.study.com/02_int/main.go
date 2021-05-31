package main

import "fmt"

func main() {
	i1 := 100
	i2 := 077
	i3 := 0x12343ef

	fmt.Printf("i1:%d\n", i1)
	fmt.Printf("i1:%b\n", i1)	// 二进制表示
	fmt.Printf("i1:%o\n", i1)	// 八进制表示
	fmt.Printf("i1:%x\n", i1)	// 十六进制表示
	fmt.Printf("i1:%T\n", i1)	// 表示类型
	fmt.Printf("i1:%v\n", i1)	// 表示值
	fmt.Println("====================")

	fmt.Printf("i2:%d\n", i2)
	fmt.Printf("i2:%o\n", i2)
	fmt.Printf("i2:%T\n", i2)
	fmt.Println("====================")

	fmt.Printf("i3:%d\n", i3)
	fmt.Printf("i3:%x\n", i3)
	fmt.Printf("i3:%T\n", i3)
	fmt.Println("====================")

	i4 := int8(67)
	fmt.Printf("i4:%d\n", i4)
	fmt.Printf("i4:%o\n", i4)
	fmt.Printf("i4:%T\n", i4)
}
