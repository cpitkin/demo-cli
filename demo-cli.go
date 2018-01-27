package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	strParams   = flag.NewFlagSet("int", flag.ExitOnError)
	intParams   = flag.NewFlagSet("string", flag.ExitOnError)
	inputIteger = intParams.Int("i", 0, "Add an integer please.")
	inputString = strParams.String("s", "", "Add a string please.")
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("One string or integer is required.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "int":
		intParams.Parse(os.Args[2:])
	case "string":
		strParams.Parse(os.Args[2:])
	}

	if intParams.Parsed() {
		if *inputIteger == 0 {
			intParams.PrintDefaults()
			os.Exit(1)
		}
	}
	if strParams.Parsed() {
		if *inputString == "" {
			strParams.PrintDefaults()
			os.Exit(1)
		}
	}

	fmt.Printf("String: %v Integer: %v", *inputString, *inputIteger)
}
