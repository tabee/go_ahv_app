package main

import (
	"fmt"

	c "github.com/tabee/go-ahv-app/pkg/ahvvalidieren"
)

func main() {
	fmt.Println("hello world")
	arr := [2]string{"56.9217.0769.84", "756.3903.3333.83"}
	for j := 0; j < len(arr); j++ {
		v, _ := c.Validate(arr[j])
		println(v)
	}
}
