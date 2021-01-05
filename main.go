package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
)

const version = "0.2.0"

type argConfig struct {
	version   bool
	separator string

	// args are the positional (non-flag) command-line arguments.
	args []string
}

// parseFlags parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.args[1:] as 'args'.
// Returns the Config in case parsing succeeded, or an error. In any case, the
// output of the flag.Parse is returned in output.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
func parseFlags(progname string, args []string) (config *argConfig, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf argConfig

	flags.BoolVar(&conf.version, "v", false, "Prints current reverb version")
	flags.StringVar(&conf.separator, "d", "-", "The character used to draw the full-width separator")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	conf.args = flags.Args()
	return &conf, buf.String(), nil
}

func reverb(width int, conf *argConfig, writer io.Writer) {
	if conf.version {
		fmt.Fprintf(writer, version+"\n")
		return
	}

	fmt.Fprintf(writer, strings.Repeat(conf.separator, width)+"\n")

	if len(conf.args) > 0 {
		fmt.Fprintf(writer, strings.Join(conf.args, " ")+"\n")
		fmt.Fprintf(writer, strings.Repeat(conf.separator, width)+"\n")
	}
}

func main() {
	conf, output, err := parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:\n", output)
		os.Exit(1)
	}

	cols, _ := consolesize.GetConsoleSize()

	reverb(cols, conf, os.Stdout)
}
