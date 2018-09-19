package main

import (
	"fmt"
	"os"
	"os/exec"
)

func depgraph() {
	const (
		godepgraphCmdName = "godepgraph"
		dotCmdName        = "dot"
		target            = "github.com/thegtproject/gravity"
		outfile           = "../out/depgraph.png"
	)

	fmt.Println("godepgraph running...")

	checkCommand(godepgraphCmdName)
	checkCommand(dotCmdName)

	cmdgodepgraph := exec.Command(godepgraphCmdName, target)

	var godepgraphoutput []byte
	var err error
	if godepgraphoutput, err = cmdgodepgraph.CombinedOutput(); err != nil {
		fmt.Println(godepgraphCmdName, "errored with:\n", err)
		os.Exit(1)
	}

	f, err2 := os.Create("temp.txt")
	if err2 != nil {
		fmt.Println("unable to create file:", err)
		os.Exit(1)
	}

	defer func() {
		f.Close()
		err = os.Remove("temp.txt")
		if err != nil {
			fmt.Println("unable to remove file:", err)
		}
	}()

	f.Write(godepgraphoutput)

	cmddot := exec.Command(dotCmdName, "-Tpng", "-o"+outfile, "temp.txt")

	if log, err := cmddot.CombinedOutput(); err != nil {
		fmt.Println(dotCmdName, "errored with:\n", err, "\n", string(log))
		os.Exit(1)
	}
}
