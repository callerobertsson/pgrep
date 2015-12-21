package main

// pgrep - simpel parallel grep

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Matcher struct {
	kind    MatcherKind
	pattern string
	count   int
}

type MatcherKind int

const (
	PrintMatcher = iota
	CountMatcher
	PrintCountMatcher
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

	// Primus motor
	err = printAndCount(reader, matchers)
	if err != nil {
		fmt.Printf("Error when matching: %v\n", err.Error())
		return
	}

	// Present summary
	printCounts(matchers)

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
func readArguments() (io.Reader, []Matcher, error) {

	if len(os.Args) < 3 {
		return nil, []Matcher{}, errors.New("Too few args. Must have atleast one pattern")
	}

	reader, err := getReader()
	if err != nil {
		return nil, []Matcher{}, err
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
func getMatchers() ([]Matcher, error) {

	ms := []Matcher{}

	for i := 1; i < len(os.Args)-1; i += 2 {
		k, err := flagToMatcherKind(os.Args[i])
		if err != nil {
			return []Matcher{}, err
		}

		ms = append(ms, Matcher{
			kind:    k,
			pattern: os.Args[i+1],
			count:   0,
		})
	}

	return ms, nil
}

// Counts all matches and prints matches for matchers af printing kind
func printAndCount(r io.Reader, ms []Matcher) error {
	s := bufio.NewScanner(r)

	// Scan reader line by line
	for s.Scan() {
		line := s.Text()

		// Loop all matchers to count and maybe print line
		for i := 0; i < len(ms); i++ {
			match, err := regexp.Match(ms[i].pattern, []byte(line))
			if err != nil {
				return err
			}

			if match {
				ms[i].count++

				// Print if that kind of matcher
				if ms[i].kind == PrintMatcher || ms[i].kind == PrintCountMatcher {
					fmt.Printf("<%d> %v\n", i, line)
				}
			}
		}
	}

	return nil
}

// Print the counts of all counting matchers
func printCounts(ms []Matcher) {

	for i, m := range ms {
		if m.kind == CountMatcher || m.kind == PrintCountMatcher {
			fmt.Printf("Matcher %q <%d> got %d matches\n", m.pattern, i, m.count)
		}
	}

}

// Converts a cammand line flag to matcher kind
func flagToMatcherKind(flag string) (k MatcherKind, err error) {

	switch flag {
	case "-p":
		return PrintMatcher, nil
	case "-c":
		return CountMatcher, nil
	case "-pc":
		return PrintCountMatcher, nil
	}

	return k, errors.New("Unknown flag: " + flag)
}
