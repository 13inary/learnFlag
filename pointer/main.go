package main

import (
	"flag"
	"fmt"
)

// package main

var num *int

// init()
func init() {
	// can't use tmp := *(flag.Int(...))
	num = flag.Int("i", 100, "user -i to set num (default:100)")
	flag.Parse()
}

// main
func main() {
	fmt.Printf("\n========== *num >>>>>>>>>>\n*num = %+v\n========== *num <<<<<<<<<<\n\n", *num)
}
