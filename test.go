package main

import (
	"fmt"
)

func main() {
	a := "abcddd"
	for _, v := range a {
		fmt.Println(string(v))
	}
}
