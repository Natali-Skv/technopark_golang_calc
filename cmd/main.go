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
	if err == nil {
		fmt.Printf("%v\n", val)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "e"},
		{Text: "pi"},
		{Text: "phi"},
		{Text: "sqrt2"},
		{Text: "sqrte"},
		{Text: "sqrtpi"},
		{Text: "sqrtphi"},
		{Text: "ln2"},
		{Text: "log2e"},
		{Text: "ln10"},
		{Text: "log10e"},
		{Text: "abs"},
		{Text: "acos"},
		{Text: "acosh"},
		{Text: "asin"},
		{Text: "asinh"},
		{Text: "atan"},
		{Text: "atan2"},
		{Text: "atanh"},
		{Text: "cbrt"},
		{Text: "ceil"},
		{Text: "copysign"},
		{Text: "cos"},
		{Text: "cosh"},
		{Text: "dim"},
		{Text: "erf"},
		{Text: "erfc"},
		{Text: "erfcinv"},
		{Text: "erfinv"},
		{Text: "exp"},
		{Text: "exp2"},
		{Text: "expm1"},
		{Text: "floor"},
		{Text: "gamma"},
		{Text: "hypot"},
		{Text: "j0"},
		{Text: "j1"},
		{Text: "log"},
		{Text: "log10"},
		{Text: "log1p"},
		{Text: "log2"},
		{Text: "logb"},
		{Text: "max"},
		{Text: "min"},
		{Text: "mod"},
		{Text: "nan"},
		{Text: "nextafter"},
		{Text: "pow"},
		{Text: "remainder"},
		{Text: "round"},
		{Text: "roundtoeven"},
		{Text: "sin"},
		{Text: "sinh"},
		{Text: "sqrt"},
		{Text: "tan"},
		{Text: "tanh"},
		{Text: "trunc"},
		{Text: "y0"},
		{Text: "y1"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("calculator>> "),
	)
	p.Run()
}
