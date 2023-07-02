package main

import (
	"fmt"
	"log"

	pronto "github.com/na4ma4/go-prontohex"
)

func main() {
	code := "0000 0067 0000 0015 0060 0018 0018 0018 0030 0018 0030 0018 0030 0018 0018 " +
		"0018 0030 0018 0018 0018 0018 0018 0030 0018 0018 0018 0030 0018 0030 0018 0030 0018 0018 0018 " +
		"0018 0018 0030 0018 0018 0018 0018 0018 0030 0018 0018 03f6"

	pc, err := pronto.FromString(code)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(pc.InfoDump())
}
