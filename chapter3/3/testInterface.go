package main

import (
	"fmt"
)

func main() {
	fmt.Print("666666666666")

	var file1 IFile = new(File)
	var file2 IRead = new(File)
	var file3 IClose = new(File)
	file1.Close()
	byte1 := []byte{}
	file2.Read(byte1)
	file3.Close()

	// 接口赋值
	var a Integer = 1
	// 可以赋值
	var b LessAdder = &a
	//报错
	// var c LessAdder = a
	fmt.Println(b)
	// fmt.Println(c)

	// Any 类型，由于Go语言中任何对象实例都满足空接口interface{}，所以interface{}看起来像是可以指向任何对象的Any类型，
	var v1 interface{} = 1
	var v2 interface{} = "abc"
	var v3 interface{} = &v2
	var v4 interface{} = struct{ X int }{1}
	var v5 interface{} = &struct{ X int }{1}
	fmt.Println(v1, v2, v3, v4, v5)
}

// 定义一个 File 结构体, File 不知道下面的 IFile， IClose 这些接口，但是他实现了这些接口，可以进行赋值
type File struct {
	name string
}

func (file *File) Read(buf []byte) (n int, err error) {
	return 0, nil
}
func (file *File) Close() (err error) {
	return nil
}

// 定义接口
type IFile interface {
	Read(buf []byte) (n int, err error)
	Close() (err error)
}

type IRead interface {
	Read(buf []byte) (n int, err error)
}

type IClose interface {
	Close() (err error)
}

// Integer 类型
type Integer int

// *Integer 会自动生成 Less() 这个方法
func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

// LessAdder 接口
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Stringer interface {
	String() string
}

// func Println(args ...interface{}) {
// 	for _, arg := range args {
// 		switch v := v1.(type) {
// 		case int:
// 		case string:
// 		default:
// 			if v, ok := arg.(Stringer); ok {

// 			} else {

// 			}
// 		}
// 	}
// }
