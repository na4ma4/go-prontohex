package main

import (
	"flag"
	"fmt"
	"os"

	pronto "github.com/na4ma4/go-prontohex"
)

//nolint:gochecknoglobals // usage for diff command.
var diffUsage = `Diff two supplied prontohex codes.

Usage: prontohex info <codeA> <codeB>
`

//nolint:forbidigo,gochecknoglobals // CLI tool output.
var diffFunc = func(cmd *Command, args []string) {
	if len(args) < 2 { //nolint:gomnd // diff between _TWO_ codes.
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

	diffHeader(
		"Source",
		(phLeft.GetSource() == phRight.GetSource()),
		phLeft.GetSource(),
		phRight.GetSource(),
	)

	diffHeader(
		"Frequency",
		phLeft.GetFrequency().Equal(phRight.GetFrequency()),
		uint16(phLeft.GetFrequency()),
		uint16(phRight.GetFrequency()),
	)

	diffBurstPairSequences("Burst Pair Sequence A", phLeft.GetBurstPairA(), phRight.GetBurstPairA())

	diffBurstPairSequences("Burst Pair Sequence B", phLeft.GetBurstPairB(), phRight.GetBurstPairB())

	os.Exit(0)
}

func NewDiffCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("diff", flag.ExitOnError),
		Execute: diffFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, diffUsage)
	}

	return cmd
}

//nolint:forbidigo // CLI tool output.
func diffHeader(hdrTitle string, equal bool, left, right uint16) {
	diffStr := " "
	if !equal {
		diffStr = "|"
	}
	fmt.Printf(
		"%s\n\t\t%d\t%s\t%d\n",
		hdrTitle+":",
		left,
		diffStr,
		right,
	)
}

//nolint:forbidigo // CLI tool output.
func diffBurstPairSequences(seqTitle string, left pronto.BurstPairSequence, right pronto.BurstPairSequence) {
	lenBPS := len(left)
	if len(right) > lenBPS {
		lenBPS = len(right)
	}

	fmt.Printf("%s[%d]:\n", seqTitle, lenBPS)
	for idx := 0; idx < lenBPS; idx++ {
		leftStr := "         "
		equal := false
		if leftA := left.Get(idx); leftA != nil {
			leftStr = leftA.String()
			equal = leftA.Equal(right.Get(idx))
		}

		rightStr := "         "
		if rightA := right.Get(idx); rightA != nil {
			rightStr = rightA.String()
		}

		diffChr := " "
		if !equal {
			diffChr = "|"
		}

		fmt.Printf(
			"\t%s\t%s\t%s\n",
			leftStr,
			diffChr,
			rightStr,
		)
	}
}
