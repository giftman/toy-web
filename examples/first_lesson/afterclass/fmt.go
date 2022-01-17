package main

import "fmt"

func main() {
	print(printNumWith2(0.3432))
	print(printBytes([]byte("hello")))
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	return fmt.Sprintf("%.2f", float642)
}

func printBytes(data []byte) string {
	return string(data)
}
