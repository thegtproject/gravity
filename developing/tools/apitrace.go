package main

import (
	"fmt"
	"os/exec"
)

func apitrace() {
	const (
		apitraceCmdName  = "apitrace"
		qapitraceCmdName = "qapitrace"
		tracefile        = "../out/developing.trace"
		target           = "../developing.exe"
	)

	fmt.Println("apitrace running...")
	checkCommand(apitraceCmdName)
	checkCommand(qapitraceCmdName)
	fmt.Println("starting apitrace. exit the opengl program when you're ready...")
	cmd := exec.Command(apitraceCmdName,
		"trace", "--api=gl", "--output="+tracefile, target,
	)

	if log, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(apitraceCmdName, "errored with:\n", err, "\n", string(log))
		return
	}

	if log, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(apitraceCmdName, "errored with:\n", err, "\n", string(log))
		return
	}

}
