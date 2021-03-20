package main

import (
	"fmt"

	basicOperation "github.com/mahoro1101/learn-goMath/basic"
)

func main() {
	var a, b int = 32, 30
	fmt.Println(basicOperation.Add(a, b))
	fmt.Println(basicOperation.Sub(a, b))
	fmt.Println(basicOperation.Multi(a, b))
	fmt.Println(basicOperation.Div(a, b))
}
