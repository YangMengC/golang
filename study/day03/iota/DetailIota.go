package main

import "fmt"

const (
	n = 12 //此时iota已经开始计数
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println("n=", n)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
