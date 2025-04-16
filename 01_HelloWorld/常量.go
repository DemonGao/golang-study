package main

import "fmt"

func main() {
	const aa string = "aaaa"

	fmt.Println("aa", aa)

	aa = "vbbbb"

	fmt.Println("aa", aa)

	/**
	常量只能使用基本数据类型
	*/
	const b = "aaa"
	const bb string = "aaa"
	const bbb, ccc = 2, "Hello World"
	const e, f bool = true, false

	const (
		h    byte = 3
		i         = "value"
		j, k      = "v", 4
		l, m      = 5, false
	)

	const (
		Male   = "Male"
		Female = "Female"
	)

	type Gender string

	const (
		Male   Gender = "Male"
		Female Gender = "Female"
	)
}
