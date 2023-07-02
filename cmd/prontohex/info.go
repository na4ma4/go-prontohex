package main

import (
	"flag"
	"fmt"
	"os"

	pronto "github.com/na4ma4/go-prontohex"
)

//nolint:gochecknoglobals // usage for info command.
var infoUsage = `Print the information about the supplied prontohex code.

Usage: prontohex info <code>
`

//nolint:forbidigo,gochecknoglobals // CLI tool output.
var infoFunc = func(cmd *Command, args []string) {
	if len(args) < 1 {
		usageAndExit("code required")
		os.Exit(1)
	}

	ph, err := pronto.FromString(args[0])
	if err != nil {
		usageAndExit(fmt.Sprintf("error: %s", err))
	}

	sonyButton := pronto.SonyRemoteButton(ph)
	fmt.Printf("Sony Button: %d\n", sonyButton)
	sonyDeviceCode := pronto.SonyDeviceCode(ph)
	fmt.Printf("Sony Device Code: %d\n", sonyDeviceCode)

	if _, err = os.Stdout.Write([]byte(ph.InfoDump())); err != nil {
		usageAndExit(fmt.Sprintf("error: %s", err))
	}

	os.Exit(0)
}

func NewInfoCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("info", flag.ExitOnError),
		Execute: infoFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, infoUsage)
	}

	return cmd
}
