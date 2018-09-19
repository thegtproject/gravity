package main

import (
	"fmt"
	"os"
	"os/exec"
)

func checkCommand(cmd string) {
	fmt.Print("checking for tool \"" + cmd + "\"... ")
	if !cmdAvailable(cmd) {
		fmt.Println("NOT FOUND")
		fmt.Printf("error: command \"%s\" was not found", cmd)
		os.Exit(1)
	}
	fmt.Println("found")
}

func cmdAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	return true
}
