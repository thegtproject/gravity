package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func apitrace() {
	var (
		apitraceCmdName  = "apitrace"
		qapitraceCmdName = "qapitrace"
		tracefile        = filepath.Join(gravitypath, "developing", "out", "developing.trace")
		target           = filepath.Join(gravitypath, "developing", "developing")
	)

	fmt.Println("apitrace running...")
	checkCommand(apitraceCmdName)
	checkCommand(qapitraceCmdName)
	fmt.Println("starting apitrace. exit the opengl program when you're ready...")
	{ // apitrace
		cmdArgs := []string{"trace", "--api=gl", "--output=" + tracefile, target}
		cmd := exec.Command(apitraceCmdName, cmdArgs...)

		if log, err := cmd.CombinedOutput(); err != nil {
			fmt.Println(apitraceCmdName, "errored with:\n", err, "\n", string(log))
			return
		}
	}
	{ // qapitrace gui
		cmdArgs := []string{tracefile}
		cmd := exec.Command(qapitraceCmdName, cmdArgs...)

		if log, err := cmd.CombinedOutput(); err != nil {
			fmt.Println(qapitraceCmdName, "errored with:\n", err, "\n", string(log))
			return
		}
	}
}
