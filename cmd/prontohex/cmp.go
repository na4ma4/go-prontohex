package main

import (
	"flag"
	"fmt"
	"os"

	pronto "github.com/na4ma4/go-prontohex"
)

//nolint:gochecknoglobals // usage for cmp command.
var cmpUsage = `Compare two supplied prontohex codes.

Usage: prontohex info <codeA> <codeB>
`

//nolint:forbidigo,gochecknoglobals // CLI tool output.
var cmpFunc = func(cmd *Command, args []string) {
	if len(args) < 2 { //nolint:gomnd // comparing _TWO_ codes.
		fmt.Println("two codes required")
		os.Exit(1)
	}

	phLeft, err := pronto.FromString(args[0])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	phRight, err := pronto.FromString(args[1])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	eqv, msg := phLeft.EqualDetail(phRight)

	fmt.Printf("Code Equality: %t\nMessage: %s\n", eqv, msg)

	os.Exit(0)
}

func NewCompareCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("cmp", flag.ExitOnError),
		Execute: cmpFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, cmpUsage)
	}

	return cmd
}
