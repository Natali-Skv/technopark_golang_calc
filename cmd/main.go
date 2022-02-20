package main

import (
	"calculator"
	"fmt"
	"os"
)

func main() {
	val, err := calculator.Calculate("((((1)+1))*(1-1)*3-(1+1))")
	if err == nil {
		fmt.Printf("%v\n", val)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// import (
// 	"calculator"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/c-bata/go-prompt"
// )

// func executor(s string) {
// 	// fmt.Print(s)
// 	s = strings.TrimSpace(s)
// 	if s == "" {
// 		return
// 	}
// 	if s == "exit" || s == "quit" {
// 		os.Exit(0)
// 	}
// 	val, err := calculator.Calculate(s)
// 	if err == nil {
// 		fmt.Printf("%v\n", val)
// 	} else {
// 		fmt.Fprintf(os.Stderr, "%v\n", err)
// 	}
// }

// func main() {
// 	p := prompt.New(
// 		executor,
// 		// TODO: add autocomplite for functions
// 		func(d prompt.Document) []prompt.Suggest { return []prompt.Suggest{} },
// 		prompt.OptionPrefix("calculator> "),
// 	)
// 	p.Run()
// }
