package main

import (
	"fmt"
	"os"
)

var gravitypath string

func main() {
	fmt.Println("Gravity Development Toolkit")

	path, err := resolveGravityPath()
	if err != nil {
		fmt.Println("unable to resolve gravity directory:", err)
		os.Exit(1)
	}
	gravitypath = path

	if !(len(os.Args) > 1) {
		fmt.Println("no command specified")
		os.Exit(1)
	}

	switch cmd := os.Args[1]; cmd {
	case "apitrace":
		apitrace()
	case "rundevapp":
		rundevapp()
	case "depgraph":
		depgraph()
	case "genversion":
		genversion()
	default:
		fmt.Println("command", cmd, "unknown")
		os.Exit(1)
	}
}
