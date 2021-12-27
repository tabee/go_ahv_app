package main

import (
	"fmt"

	"github.com/tabee/go-ahv-app/pkg/ahvvalidieren"
)

// demo how to use
func main() {
	fmt.Println("hello world")
	arr := [2]string{"756.9217.0769.84", "756.3903.3333.83"}
	for j := 0; j < len(arr); j++ {
		v, _ := ahvvalidieren.Validate(arr[j])
		println(v)
	}
}
