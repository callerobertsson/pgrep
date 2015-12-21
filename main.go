// pgrep - simpel parallel grep

package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"./matcher"
)

func printUsage() {
	fmt.Println(`
PGREP - Parallel grep tool

usage: pgrep --help
   or: pgrep <flag> <pattern> [<flag> <pattern> ...] [file]      

flag: -p       print matching lines
  or: -c       count matching lines
  or: -pc      print and count matching lines

pattern: golang regex

example: pgrep -p 'foo' -c '^\s*$' -pc ';$' bar.txt

    print all lines containing foo
	count all empty lines
	print and count all lines ending with ;
	`)
}

func main() {

	if len(os.Args) < 2 || hasHelpArgument() {
		printUsage()
		return
	}

	// Parse command line
	reader, matchers, err := readArguments()
	if err != nil {
		fmt.Printf("Error reading arguments: %v\n", err.Error())
		return
	}

	// Create matcher
	m := matcher.New(matchers, reader, os.Stdout)

	// Primus motor
	err = m.PrintAndCount()
	if err != nil {
		fmt.Printf("Error when matching: %v\n", err.Error())
		return
	}

	// Present summary
	m.PrintCounts()

}

// Check if first argument is a help flag
func hasHelpArgument() bool {

	if len(os.Args) < 2 {
		// no args => no help flag
		return false
	}

	switch os.Args[1] {
	case "-h", "--help", "-?", "--about":
		return true
	}

	return false
}

// Read arguments to get matchers and reader
func readArguments() (io.Reader, []matcher.Match, error) {

	if len(os.Args) < 3 {
		return nil, []matcher.Match{}, errors.New("Too few args. Must have atleast one pattern")
	}

	reader, err := getReader()
	if err != nil {
		return nil, []matcher.Match{}, err
	}

	ms, err := getMatchers()

	return reader, ms, err
}

// Get reader from file or stdin
func getReader() (reader io.Reader, err error) {

	if len(os.Args)%2 == 0 {
		// last arg is file
		file := os.Args[len(os.Args)-1]
		reader, err = os.Open(file)
	} else {
		// no file last in arguments
		reader = os.Stdin
	}

	return
}

// Get matchers from cammand line
func getMatchers() ([]matcher.Match, error) {

	ms := []matcher.Match{}

	for i := 1; i < len(os.Args)-1; i += 2 {
		k, err := matcher.FlagToMatchKind(os.Args[i])
		if err != nil {
			return []matcher.Match{}, err
		}

		ms = append(ms, matcher.Match{
			Kind:    k,
			Pattern: os.Args[i+1],
			Count:   0,
		})
	}

	return ms, nil
}
