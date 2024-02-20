package main

import (
	"flag"
	"fmt"

	"github.com/integralist/test-go-binary/internal/flags"
)

func main() {
	var msg, warning string
	flag.StringVar(&msg, "message"+flags.MessageVersion, "", "message to be displayed")
	flag.StringVar(&warning, "warning", "", "warning to be displayed")
	flag.Parse()
	fmt.Printf("message: %+v\n", msg)
	fmt.Printf("warning: %+v\n", warning)
}
