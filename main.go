package main

import (
	"flag"
	"fmt"
)

func main() {
	var msg, warning string
	flag.StringVar(&msg, "messagev3", "", "message to be displayed")
	flag.StringVar(&warning, "warning", "", "warning to be displayed")
	flag.Parse()
	fmt.Printf("message: %+v\n", msg)
	fmt.Printf("warning: %+v\n", warning)
}
