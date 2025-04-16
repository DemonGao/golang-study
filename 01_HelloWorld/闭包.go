package main

import "fmt"

func main() {
	/**
	闭包
	*/
	a := 10
	total := func(a int, b int) int {
		fmt.Println(a, b)
		return a + b
	}(a, 4)
	fmt.Println(total)

	fn := func() func(a int, b int) int {
		var x int = 1
		return func(a int, b int) int {
			x++
			fmt.Println(a, b)
			return x
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(fn(i, 8))
	}
}
