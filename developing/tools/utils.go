package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func resolveGravityPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("unable to resolve Gravity path:", err)
		os.Exit(1)
	}
	gravityPath := filepath.Join(filepath.Dir(cwd), "gravity")
	if exists, err := dirExists(gravityPath); !exists {
		return "", err
	}
	return gravityPath, nil
}

func dirExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}
	return true, nil
}
