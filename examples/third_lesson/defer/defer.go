package main

import "fmt"

//相当于 finally 一般用来释放资源
//先执行的后执行
func main() {
	defer func() {
		fmt.Println("aaa")
	}()

	defer func() {
		fmt.Println("bbb")
	}()

	defer func() {
		fmt.Println("ccc")
	}()
}
