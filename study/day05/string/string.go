package main

import (
	"fmt"
	"strings"
)

func main()  {
	s1 := `
    世情薄，
    人情恶，
    `
	s2 := "雨送黄昏花易落。"
	s3 := "higoh,agiaohg,agniahog,agiahog"
	fmt.Printf("%s\n",s1+s2)
	fmt.Print(len(s1))
	ss := fmt.Sprintf("%s%s",s1,s2)
	//s1 = ss
	fmt.Println(ss)
	ret := strings.Split(s3,",")
	fmt.Println(ret[0])
}
