package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i1 int = 1
	var s1 string = "s1"
	fmt.Println(i1, s1)

	/**
	TODO: 1.4.1 指针声明与初始化
	*/
	// var p1 *int = &i1
	var p1 *int
	p1 = &i1

	// var p2 *string = &s1
	var p2 *string
	p2 = &s1
	p3 := &s1

	var p4 *int // 指针的零值 = <nil>

	var pp1 **int = &p1
	fmt.Println("输出指针", p1, p2, p3, p4, pp1)

	fmt.Println(*p1, *p2)

	/**
	TODO: 1.4.2 使用指针访问值
	使用 * 引用指针来访问指针指向的值
	*/
	println("使用指针访问值", *p1, *p2, **pp1)

	/**
	TODO: 1.4.3 修改指针指向的值
	*/
	a := 2
	p1 = &a
	a = 4
	fmt.Println(*p1, p1, &a)

	pp1 = &p1
	fmt.Println(pp1, &p1)
	**pp1 = 3

	fmt.Println(pp1, a)
	//var pp1 **int = &p1

	/**
	TODO: 1.4.4 指针、unsafe.Pointer 和 uintptr
	但是 Go 中提供了其他方式，来操作指针，即引入了 unsafe.Point 类型和 uintptr 类型，来帮助我们操作指针。
	uintptr 类型是把内存地址转换成了一个整数，然后对这个整数进行计算后，在把 uintptr 转换成指针，达到指针偏移的效果。
	unsafe.Pointer 是普通指针与 uintptr 之间的桥梁，通过 unsafe.Pointer 实现三者的相互转换。
	*/
	var aa string = "Hellow, world!"
	upA := uintptr(unsafe.Pointer(&aa))
	upA += 1

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)
}
