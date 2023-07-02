package main

import (
	"flag"
	"fmt"
	"os"
)

//nolint:gochecknoglobals // usage for command.
var usage = `Usage: prontohex command [options]

A simple tool to parse and manipulate prontohex codes

Options:

Commands:
  cmp		Compares two prontohex commands
  diff		Displays a diff of two prontohex commands
  info		Displays information about supplied pronto hex code
  version	Prints version info to console
`

//nolint:forbidigo // CLI tool output.
func main() {
	var cmd *Command

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	if len(os.Args) <= 1 {
		usageAndExit("")
	}

	switch os.Args[1] {
	case "version":
		cmd = NewVersionCommand()
	case "info":
		cmd = NewInfoCommand()
	case "cmp":
		cmd = NewCompareCommand()
	case "diff":
		cmd = NewDiffCommand()
	default:
		usageAndExit(fmt.Sprintf("%s: '%s' is not a command.\n", os.Args[0], os.Args[1]))
	}

	if err := cmd.Init(os.Args[2:]); err != nil {
		fmt.Printf("error: command init failed: %s", err)
		os.Exit(1)
	}

	cmd.Run()
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}
