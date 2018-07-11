package main

import "fmt"
import "unsafe"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("The area is : %d", area)
	println()
	println(a, b, c)
}
