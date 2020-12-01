package main

import (
	"fmt"
	"unicode"
)

func main()  {
	s := "你说hello我说mengc"
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			fmt.Printf(string(c))
		}
	}
}
