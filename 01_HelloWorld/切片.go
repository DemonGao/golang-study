package main

import "fmt"

func main0416() {

	// 数组定义 var 数组名 [元素个数]数据类型
	var num [3]int = [3]int{1, 2, 3}
	fmt.Println("num = ", num)

	// 切片定义  var 切片名 []数据类型
	//var s []int
	//fmt.Println(s)

	// 自动推导类型创建
	s := make([]int, 5)
	s[0] = 123
	s[1] = 234
	s[2] = 345
	s[3] = 456
	s[4] = 567
	//s[5] = 678

	fmt.Println(s)

	// 通过 length 查看切片长度
	fmt.Println(len(s))

	s = append(s, 789, 789, 8910)
	fmt.Println(len(s))
}

func main0416_1() {
	s := make([]int, 5)
	s[0] = 123
	s[1] = 234
	s[2] = 345
	s[3] = 456
	s[4] = 567

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}

}

func main0416_2() {
	var s []int = []int{1, 2, 3, 4, 5}

	s = append(s, 6, 7, 8, 9)
	fmt.Println("长度: ", len(s))
	fmt.Println("容量: ", cap(s))

	// 容量大于等于长度
	s = append(s, 6, 7, 8, 9)
	fmt.Println("长度: ", len(s))
	fmt.Println("容量: ", cap(s))

	// 容量每次扩展为上次的倍数,
}

func main0416_3() {
	var s []int = []int{1, 2, 3, 4, 5}

	s = append(s, 6, 7, 8, 9)
	fmt.Println("长度: ", len(s))
	fmt.Println("容量: ", cap(s))

	// 容量大于等于长度
	s = append(s, 6, 7, 8, 9)
	fmt.Println("长度: ", len(s))
	fmt.Println("容量: ", cap(s))

	// 容量每次扩展为上次的倍数, 超过 1024个时，增长倍数为 1/4

	s1 := s[3:]
	fmt.Println(s1)
}

func main0416_4() {
	var s []int = []int{1, 2, 3, 4, 5}

	s = append(s, 6, 7, 8, 9)
	fmt.Println("长度: ", len(s))
	fmt.Println("容量: ", cap(s))

	// 容量每次扩展为上次的倍数, 超过 1024个时，增长倍数为 1/4

	s1 := s[3:]
	fmt.Println(s1)

	s2 := s[0:3]
	fmt.Println(s2)

	s3 := s[:3]
	fmt.Println(s3)
}

// 当切片是基于同一个数组指针创建出来时，修改数组中的值时，同样会影响到这些切片。
func main0416_5() {
	a := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := a[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := a[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	a[0] = 9
	a[1] = 8
	a[2] = 7
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
}

func main0416_6() {
	s1 := []int{5, 4, 3, 2, 1}
	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Println(s1)
	fmt.Println(e1, e2, e3)

	// 向指定位置赋值
	s1[0] = 10
	s1[1] = 9
	s1[2] = 8
	fmt.Println(s1)
	fmt.Println(e1, e2, e3)

	e4 := s1[3:4]
	e4[0] = 7
	fmt.Println(s1)
	fmt.Println(e4, s1[3])

	// range迭代访问切片
	for i, v := range s1 {
		fmt.Printf("before modify, s1[%d] = %d\n", i, v)
	}
}

func main0416_7() {
	s3 := []int{}
	fmt.Println("s3 = ", s3)

	// append函数追加元素
	s3 = append(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	fmt.Println("s3 = ", s3)

	// 向指定位置添加元素
	s4 := []int{1, 2, 4, 5}
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...) // [1,2,3,4,5]
	fmt.Println("s4 = ", s4)

	// 移除指定位置元素代码
	s5 := []int{1, 2, 3, 4, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("s5 = ", s5)

}

func main() {
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)
}
