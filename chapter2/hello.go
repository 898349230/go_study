package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hello")
	var v1 = 1
	var v2 int = 2
	var v3 string

	fmt.Println(v1, v2, v3)

	v4 := 10
	fmt.Println("v4 = ", v4)

	var b, c int = 1, 3
	fmt.Println(b, c)

	// 匿名变量
	_, _, nickName := GetName()
	fmt.Println("nickName : ", nickName)

	fmt.Println(Pi)
	fmt.Println("size ", size, " eof", eof)

	fmt.Println("c0 ", c0, " c1 ", c1, " d0 ", d0, " d1", d1)

	var hello = "hello, 世界"
	char := hello[0]
	fmt.Print("  ", hello, " ", char, " length ", len(hello))

	for i, ch := range hello {
		fmt.Println("遍历字符串 ", i, " ", ch)
	}

	fmt.Println("********** 切片 *******")
	var myArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySlice := myArray[:5]

	fmt.Println("打印数组")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("打印切片")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	// make() 创建数组切片
	fmt.Println("make() 创建数组切片")
	mySlice1 := make([]int, 5)
	for i := 0; i < len(mySlice1); i++ {
		fmt.Print(mySlice1[i], " ")
	}
	fmt.Println("make() 创建数组切片2")
	// 初始化元素个数为5，初始值为0，并预留10个元素的的存储空间
	mySlice2 := make([]int, 5, 10)
	for i, v := range mySlice2 {
		fmt.Println("mySlice2 [", i, "] = ", v)
	}

	// cap() 数组切片的分片大小， len() 当前存储的元素个数
	fmt.Println("cap(mySlice2) ", cap(mySlice2))
	fmt.Println("len(mySlice2) ", len(mySlice2))

	mySlice = append(mySlice, 100, 200)
	mySlice = append(mySlice, mySlice2...)
	mySlice3 := mySlice[:2]
	for i, v := range mySlice3 {
		fmt.Println("mySlice3 [", i, "] = ", v)
	}

	// 复制 mySlice2 的元素到 mySlice 中，个数以 mySlice 为准
	copy(mySlice, mySlice2)

	// map
	// 声明 map
	var personDB map[string]PersonInfo
	// 通过 make() 函数创建map
	// personDB = make(map[string]PersonInfo)
	// 声明存储能力为100的map
	personDB = make(map[string]PersonInfo, 100)

	// 元素赋值
	personDB["a"] = PersonInfo{"1", "张三", "北京"}
	personDB["b"] = PersonInfo{"2", "李四", "上海"}

	person, ok := personDB["b"]

	if ok {
		fmt.Println("found person ", person.Name)
	} else {
		fmt.Println("no found ")
	}

	sum := 0
	for {
		sum++
		if sum > 100 {
			fmt.Println("sum = ", sum)
			break
		}
	}

	arr1 := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(arr1)-1; i < j; i, j = i+1, j-1 {
		arr1[i], arr1[j] = arr1[j], arr1[i]
	}
	fmt.Println("数组交换")
	for _, v := range arr1 {
		fmt.Print(v)
	}

	fmt.Println("break")

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j > 3 {
	// 		}
	// 		fmt.Print(j)
	// 	}
	// }
	fmt.Println("break 到")

	// ret, err := Add(10, 12)
	ret, err := Add(-4, 12)
	fmt.Println("ret ", ret, " error ", err)

	fmt.Println("可变参数测试 ", myfunc(1, 2, 3, 4, 5))

	// 匿名函数
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println("匿名函数调用 ", f(10, 11))

	fmt.Println("测试闭包")
	var j int = 5
	aa := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	aa()
	j *= 2
	aa()
}

// 小写字母开头的函数只在本包内可见，大写字母开头的函数才能被其他包使用
func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("a and b < 0")
		return
	}
	return a + b, nil
}

func myfunc(args ...int) (ret int) {
	result := 0
	for _, i := range args {
		result += i
	}
	return result
}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func GetName() (firstName, lastName, nickName string) {
	return "张", "三", "小张"
}

// 常量
const Pi float64 = 3.1415026
const (
	size int = 101
	eof      = -1
)

// 在下一个 const 出现之前 iota 是从0开始自增的
const (
	c0 = iota
	c1 = iota
	c2 = iota
)

// iota 重新从 0 开始自增
const (
	d0 = iota
	d1 = iota
	d2 = iota
)
