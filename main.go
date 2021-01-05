package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
)

const version = "0.1.0"

func main() {
	versionFlag := flag.Bool("v", false, "prints current roxy version")
	separatorCharacter := flag.String("d", "-", "The character used to draw the full-width separator")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	cols, _ := consolesize.GetConsoleSize()

	fmt.Println(strings.Repeat(*separatorCharacter, cols))

	if len(flag.Args()) > 0 {
		fmt.Println(flag.Args()[0])
		fmt.Println(strings.Repeat(*separatorCharacter, cols))
	}
}
