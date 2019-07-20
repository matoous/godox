package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagPrintHelp := flag.Bool("h", false, "Print help")
	flagIncludeTests := flag.Bool("t", false, "Include tests")
	// print usage
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: godox [path] [path] ...\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flagPrintHelp {
		flag.Usage()
		return
	}

	paths := flag.Args()
	// by default traverse whole tree down from the directory where godox was called
	if len(paths) == 0 {
		paths = []string{"./..."}
	}
	includeTests := *flagIncludeTests

	exitWithError := false
	for _, path := range paths {
		messages, err := godox(path, includeTests)
		for _, message := range messages {
			_, _ = fmt.Fprintf(os.Stdout, "%s\n", message)
			exitWithError = true
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
			exitWithError = true
		}
	}

	if exitWithError {
		os.Exit(1)
	}
}
