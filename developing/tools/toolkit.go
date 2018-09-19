package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Gravity Development Toolkit")

	if !(len(os.Args) > 1) {
		fmt.Println("no command specified")
		os.Exit(1)
	}

	switch cmd := os.Args[1]; cmd {
	case "apitrace":
		apitrace()
	case "depgraph":
		depgraph()
	case "genversion":
		genversion()
	default:
		fmt.Println("command", cmd, "unknown")
		os.Exit(1)
	}
}
