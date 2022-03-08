package main

import (
	"calculator"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	if s == "exit" || s == "quit" {
		os.Exit(0)
	}
	val, err := calculator.Calculate(s)
	if err == nil {
		fmt.Printf("%v\n", val)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func main() {

	flag.Parse()

	if len(flag.Args()) == 1 {
		fmt.Println(flag.Args()[0])
		executor(flag.Args()[0])
		return
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("calculator>> "),
	)
	p.Run()
}
