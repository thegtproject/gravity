package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func rundevapp() {
	var (
		runCmd     = "go"
		runCmdArgs = []string{"run", "."}
		runDir     = filepath.Join(gravitypath, "developing")
	)

	genversion()

	fmt.Println("rundevapp running...")

	checkCommand(runCmd)

	fmt.Print("------\n\n")

	cmdrun := exec.Command(runCmd, runCmdArgs...)
	cmdrun.Dir = runDir
	cmdrun.Stdout = os.Stdout
	cmdrun.Stderr = os.Stderr

	if err := cmdrun.Run(); err != nil {
		fmt.Printf("command %s failed: %s\n\n", runCmd, err)
		os.Exit(1)
	}
}
