package main

import "fmt"

var count = 0

func test() {
	if count > 10 {
		return
	}
	fmt.Println(count)
	count++
	test()
}

func main() {
	test()
}
