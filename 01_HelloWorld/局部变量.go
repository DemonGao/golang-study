package main

import (
	"fmt"
	"time"
)

func main() {
	var a int

	fmt.Println("a = ", a)

	if b := 1; b == 0 {
		fmt.Println("b === 0")
	} else {
		c := 2
		fmt.Println("declare c =", c)
		fmt.Println("b == 1")
	}

	switch d := 3; d {
	case 1:
		e := 4
		fmt.Println("declare e = ", e)
		fmt.Println("d == 1")
	case 3:
		f := 4
		fmt.Println("declare f = ", f)
		fmt.Println("d == 3")
	}

	select {
	case <-time.After(time.Second):
		selectA := 1
		fmt.Println("selectA = ", selectA)
	}
	// fmt.Println(e)
	// fmt.Println(f)

}
