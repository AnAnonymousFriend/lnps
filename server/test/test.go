package main

import (
	"fmt"
)

var T int64 = a()

func init() {
	fmt.Println("init in main.go ")
}

func a() int64 {
	fmt.Println("calling a()")
	return 2
}
func main() {
	fmt.Println("calling main")
}