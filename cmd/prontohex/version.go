package main

import (
	"flag"
	"fmt"
	"os"
)

//nolint:gochecknoglobals,unused // these have to be variables for the linker to change the values
var (
	version = "dev"
	date    = "notset"
	commit  = ""
	builtBy = ""
	repo    = ""
	short   bool
)

//nolint:gochecknoglobals // usage for version command.
var versionUsage = `Print the app version and build info for the current context.

Usage: prontohex version [options]

Options:
  --short  If true, print just the version number. Default false.
`

//nolint:forbidigo,gochecknoglobals // CLI tool output.
var versionFunc = func(cmd *Command, args []string) {
	if short {
		fmt.Printf("%s", version)
	} else {
		fmt.Printf("%s [%s] (%s) <%s>", version, commit, date, builtBy)
	}
	os.Exit(0)
}

func NewVersionCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	cmd.flags.BoolVar(&short, "short", false, "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}

	return cmd
}
