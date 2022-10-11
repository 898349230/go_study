package main

import (
	"fmt"
	// "hello/chapter3/baseOne"
)

// 为类型添加方法
// 你可以给任意类型（包括内置类型，但不包括指针类型）添加相应的方法，定义了一个新类型Integer，
// 它和int没有本质不同，只是它为内置的,int类型增加了个新方法Less()。
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println("a < 2")
	}

	var arr1 = [3]int{1, 2, 3}
	// 这里是完整的复制了数组
	var arr2 = arr1
	arr2[1] = 100
	fmt.Println(arr1, arr2)

	var arr3 = [3]int{1, 2, 3}
	// 使用指针，
	var arr4 = &arr3
	arr4[1] = 100
	fmt.Println(arr3, arr4)

	fmt.Println("结构体。。。。")

	// 初始化 Rect 对象
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}

	fmt.Println("rect1: ", rect1.Area())
	fmt.Println("rect2: ", rect2.Area())
	fmt.Println("rect3: ", rect3.Area())
	fmt.Println("rect4: ", rect4.Area())

	// 测试继承
	fmt.Println("测试继承。。。")
	base := &Base{}
	foo := &Foo{}
	base.Bar()
	foo.Bar()
	// 这里会走 Base.Foo()
	foo.Foo()

	// baseOne := &baseOne.BaseOne{}
	// baseOne.PrintStr("test")
}

// 结构体
type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

// Base 结构体， 定义了 Foo 和 Bar 方法
type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("Base.Foo()")
}
func (base *Base) Bar() {
	fmt.Println("Base.Bar()")
}

// Foo 结构体内组合了 Base ，类似继承
type Foo struct {
	Base
}

// Foo2 继承可以使用指针
type Foo2 struct {
	*Base
}

// Foo 结构体内继承并且改写了 Bar 方法
func (foo *Foo) Bar() {
	// 这里是先调用 Base.Bar()，也可以不调用
	foo.Base.Bar()
	fmt.Println("Foo.Bar()")
}
