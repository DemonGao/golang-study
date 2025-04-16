package main

import "fmt"

/**
1.5 结构体
在 Go 中，开发者可以通过类型别名(alias types)和结构体的形式支持用户自定义类型。

结构体的目的就是把数据聚合在一起，能够便捷的操作这些数据。

Go 中没有类的概念。
Go 中，全局变量、全局常量、结构体、字段、方法，只有两种公开类型，公开与非公开。非公开是针对包级别的，也就是说如果全局变量声明在不同的源文件中，但是这些源文件属于相同的包，那么这些中的代码可以引用这些不公开的全局变量。不属于相同的包就访问不到了。并且公开的属性是首字母大写，非公开的属性首字母是小写，仅按照这个规则来定义是否公开。
*/

/**
1.5.1 定义结构体
在 Go 中最常用的方法使用 type 和 struct 关键字定义一个结构体，结构体中的字段都有不同的名字，并且字段可以是任意类型，比如结构体本身、指针甚至是函数：
*/

type Other struct{}

type Person struct {
	Name string `json:"name" gorm:"column:<name>"`
	Age  int
	//Call  func()
	//Map   map[string]string
	//Ch    chan string
	//Arr   [32]uint8
	//Slice []interface{}
	//Ptr   *int
	//O     Other
}

func (person Person) getName() string {
	return person.Name
}

func (person Person) getAge() int {
	return person.Age
}

func main1() {
	var s Person
	s.Name = "张三"
	s.Age = 18

	fmt.Println(s)

	var ss Person = Person{"张三", 15}

	fmt.Println(ss)
	fmt.Println(ss.getName())
	fmt.Println(ss.getAge())

}

/*
*
1.5.1.1 定义匿名字段
结构体中的字段不是一定要有字段名，也可以仅定义类型，这种只有类型没有字段名的字段被称为匿名字段。
同一个结构体中相同类型的匿名字段只能同时存在一个，但是可以同时声明多个不同类型的匿名字段。
*/
type Custom struct {
	int
	string
	Other string
}

/**
1.5.2 定义匿名结构体
匿名结构体是没有定义名称的结构体。

匿名结构体无法定义自己的类型方法。

方式 1，仅可在函数外声明，这种方式可以看成是声明了一个匿名的结构体，实例化后赋值给了的全局变量：
*/

func NewC() Person {
	return Person{
		Name: "join",
		Age:  18,
	}
}
func main2() {
	c := NewC()
	cp := &c

	cp1 := c
	fmt.Println(c.getName())
	fmt.Println(cp.getName())
	fmt.Println(cp1.getName())

	c.Name = "waaa"
	fmt.Println(c.getName())   // waaa
	fmt.Println(cp.getName())  // waaa
	fmt.Println(cp1.getName()) // join
}

func main() {
	main1()
	main2()
}
