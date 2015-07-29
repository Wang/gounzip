package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Wang/gounzip/unzip"
)

var (
	Silence bool
)

func init() {
	flag.BoolVar(&Silence, "s", false, "is silence mode")
}

func main() {
	flag.Parse()

	var zipfile, targetDir string

	targetDir = "./" //默认在本目录
	args := flag.Args()
	if len(args) == 0 {
		Printf("Use: %s", "gounzip file.zip")
		os.Exit(1)
	} else if len(args) == 1 {
		zipfile = args[0]
	} else {
		zipfile = args[0]
		targetDir = args[1]
	}

	unzip.Silence = Silence
	if err := unzip.Do(zipfile, targetDir); err != nil {
		fmt.Errorf("unzip do: %s", err.Error())
		os.Exit(1)
	}
}

func Printf(format string, v ...interface{}) {
	if Silence {
		return
	}
	fmt.Printf(format+"\n", v...)
}
