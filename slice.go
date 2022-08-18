package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	fmt.Println("Hello", name)
}
