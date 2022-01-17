package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	fmt.Println(s)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	fmt.Println(s)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	fmt.Println(s)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	fmt.Println(s)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)
	fmt.Println(s)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)
	fmt.Println(s)

}

func Add(s []int, index int, value int) []int {
	//TODO
	s0 := s[0:index]
	s1 := []int{value}
	s2 := s[index:]
	return append(s0, append(s1, s2...)...)
}

func Delete(s []int, index int) []int {
	// TODO
	s0 := s[0:index]
	s2 := s[index+1:]
	return append(s0, s2...)
}
