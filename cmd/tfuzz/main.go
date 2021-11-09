package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/taise-hub/tfuzz"
)

func main() {
	var options tfuzz.Options
	flag.StringVar(&options.TargetUrl, "u", "", "Target URL")
	flag.StringVar(&options.InputFile, "f", "", "Input File")
	flag.Parse()

	if len(options.TargetUrl) == 0 {
		tfuzz.ShowError("Target URL(-u) is required field.\n")
		os.Exit(0)
	}

	if len(options.InputFile) == 0 {
		tfuzz.ShowError("Fuzzing File(-f) is required field.\n")
		os.Exit(0)
	}

	fmt.Println(options.TargetUrl)
	fuzzList := options.ReadFile()
	options.StartFuzz(fuzzList)
}