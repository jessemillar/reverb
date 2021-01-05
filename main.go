package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
)

func main() {
	separatorCharacter := flag.String("d", "-", "The character used to draw the full-width separator")
	flag.Parse()

	cols, _ := consolesize.GetConsoleSize()

	fmt.Println(strings.Repeat(*separatorCharacter, cols))

	if len(flag.Args()) > 0 {
		fmt.Println(flag.Args()[0])
		fmt.Println(strings.Repeat(*separatorCharacter, cols))
	}
}
