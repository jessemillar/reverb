package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
)

func main() {
	cliArgs := os.Args

	cols, _ := consolesize.GetConsoleSize()

	fmt.Println(strings.Repeat("-", cols))

	if len(cliArgs) > 1 {
		fmt.Println(cliArgs[1])
		fmt.Println(strings.Repeat("-", cols))
	}
}
