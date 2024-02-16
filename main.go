package main

import (
	"flag"
	"fmt"
)

func main() {
	var msg string
	flag.StringVar(&msg, "msg", "", "message to be displayed")
	flag.Parse()
	fmt.Printf("message: %+v\n", msg)
}
