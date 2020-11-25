package main

import "fmt"

func main() {
	var s1 = "you "
	s2 := "are good"
	//如果要修改s2的值，只能使用`s2 = 20`
	fmt.Println(s1 + s2)
}
