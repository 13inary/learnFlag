package main

import (
	"flag"
	"fmt"
)

// package main
var num int

// init
func init() {
	flag.IntVar(&num, "i", 100, "use -i to set num (default:100)")
	flag.Parse()
}

// main
func main() {
	fmt.Printf("\n========== num >>>>>>>>>>\nnum = %+v\n========== num <<<<<<<<<<\n\n", num)
}
