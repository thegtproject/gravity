package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/thegtproject/gravity"
)

var Log = gravity.Log

func main() {
	path, err := resolveBasePath()
	fmt.Println(path, err)
	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 1366, Height: 768,
		VSync: true,
	}

	gravity.Log.SetPrint(GUIOutput.Print)
	gravity.Log.SetPrintf(GUIOutput.Printf)
	gravity.Log.SetPrintln(GUIOutput.Println)
	Log = gravity.Log

	gravity.Init(cfg)
	gravity.Run(setupscene)
}

func init() {
	runtime.LockOSThread()
}

func resolveBasePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("unable to resolve base path:", err)
		os.Exit(1)
	}
	basePath := filepath.Join(filepath.Dir(cwd), "developing")
	if exists, err := dirExists(basePath); !exists {
		return "", err
	}
	return basePath, nil
}

func dirExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}
	return true, nil
}
