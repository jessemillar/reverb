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

const version = "2.2.0"
const headlessWidth = 42

type argConfig struct {
	version               bool
	disableDynamicWidth   bool
	enableEscapeSequences bool
	separator             string
	args                  []string // args are the positional (non-flag) command-line arguments
}

// parseFlags parses the command-line arguments provided to the program
// Typically os.Args[0] is provided as 'progname' and os.args[1:] as 'args'
// Returns the Config in case parsing succeeded, or an error
// In any case, the output of the flag.Parse is returned in output
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message
func parseFlags(progname string, args []string) (config *argConfig, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)

	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf argConfig

	flags.BoolVar(&conf.version, "v", false, "Print the current reverb version")
	flags.BoolVar(&conf.disableDynamicWidth, "d", false, "Disable matching the width of the separators to the length of of long strings in a headless terminal")
	flags.BoolVar(&conf.enableEscapeSequences, "e", false, "Enable parsing of escape sequences (\\, \\n, \\r, etc.)")
	flags.StringVar(&conf.separator, "c", "-", "The character used to draw the separator")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}

	conf.args = flags.Args()

	return &conf, buf.String(), nil
}

func findLongestLine(allLines []string) int {
	longest := 0

	for _, line := range allLines {
		if len(line) > longest {
			longest = len(line)
		}
	}

	return longest
}

func reverb(width int, conf *argConfig, writer io.Writer) {
	// Print the version number and exit if that's what's asked for
	if conf.version {
		fmt.Fprintf(writer, version+"\n")
		return
	}

	// Limit the separator to one character
	if len(conf.separator) > 1 {
		fmt.Fprintf(writer, "Please pass only one character to the -c flag\n")
		return
	}

	// Combine multiple arguments into a single word respecting newlines
	reverbString := strings.Join(conf.args, " ")

	// Unescape backslash characters if allowed
	if conf.enableEscapeSequences {
		reverbString = strings.Replace(reverbString, "\\a", "\a", -1)
		reverbString = strings.Replace(reverbString, "\\b", "\b", -1)
		reverbString = strings.Replace(reverbString, "\\f", "\f", -1)
		reverbString = strings.Replace(reverbString, "\\n", "\n", -1)
		reverbString = strings.Replace(reverbString, "\\r", "\r", -1)
		reverbString = strings.Replace(reverbString, "\\t", "\t", -1)
		reverbString = strings.Replace(reverbString, "\\v", "\v", -1)
	}

	// React if we're in a headless terminal
	if width == 0 {
		stringWidth := len(reverbString)

		if stringWidth > 0 && !conf.disableDynamicWidth {
			// If the passed string has multiple lines, find the width of the longest line
			if conf.enableEscapeSequences {
				stringWidth = findLongestLine(strings.Split(reverbString, "\n"))
			}

			// Take a look at the string width
			if stringWidth < headlessWidth {
				width = headlessWidth
			} else {
				// If there's a string longer than headlessWidth, match the separator width to the long string
				width = stringWidth
			}
		} else {
			// Use a static separator width
			width = headlessWidth
		}
	}

	// Print the first separator
	fmt.Fprintf(writer, strings.Repeat(conf.separator, width)+"\n")

	// Print the text and a second separator if we supplied a text string
	if len(reverbString) > 0 {
		fmt.Fprintf(writer, reverbString+"\n")
		fmt.Fprintf(writer, strings.Repeat(conf.separator, width)+"\n")
	}
}

func main() {
	// Parse command line flags
	conf, output, err := parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:\n", output)
		os.Exit(1)
	}

	// Get terminal width
	cols, _ := consolesize.GetConsoleSize()

	// Disable dynamic separator width if told to by the user
	if conf.disableDynamicWidth {
		cols = 0
	}

	// Do the magic
	reverb(cols, conf, os.Stdout)
}
