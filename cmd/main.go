package main

import (
	"calculator"
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Printf("%v\n", val)
}

func main() {
	p := prompt.New(
		executor,
		func(d prompt.Document) []prompt.Suggest { return []prompt.Suggest{} },
		prompt.OptionPrefix("calculator>> "),
	)
	p.Run()
}
