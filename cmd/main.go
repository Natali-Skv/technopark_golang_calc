package main

// import (
// 	"calculator"
// 	"fmt"
// )

// func main() {
// 	// val, err := calculator.Calculate("-1")
// 	// expr, poz, err := calculator.ScanExpression([]rune("pow((2)),1)"), 0)
// 	// expr, poz, err := calculator.ScanExpression([]rune("(1+1))"), 0)
// 	expr, err := calculator.Calculate("pow(2,10)+pow(2,3)")
// 	fmt.Println(expr)
// 	// fmt.Println(poz)
// 	fmt.Println(err)
// }

import (
	"calculator"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func executor(s string) {
	// fmt.Print(s)
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
		// {Text: "users", Description: "Store the username and age"},
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
