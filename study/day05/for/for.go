package main

import "fmt"

func main()  {
	////基本格式
	//for i:=0; i<10;i++{
	//	fmt.Println(i)
	//}
	////变种1
	//var i = 5
	//for  ;i<10;i++{
	//	fmt.Println(i)
	//}
	////变种2
	//for i<10{
	//	fmt.Println(i)
	//	i++
	//}
	s := "Hello沙河"
	for i, v := range s{
		fmt.Println("%d %c\n",i,v)
	}
}