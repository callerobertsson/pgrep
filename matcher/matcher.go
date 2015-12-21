// Package matcher

package matcher

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
)

type Matcher struct {
	matches []Match
	reader  io.Reader
	writer  io.Writer
}

type Match struct {
	Kind    MatchKind
	Pattern string
	Count   int
}

type MatchKind int

const (
	PrintMatch = iota
	CountMatch
	PrintCountMatch
)

// Create a new matcher
func New(matches []Match, reader io.Reader, writer io.Writer) Matcher {
	return Matcher{
		matches: matches,
		reader:  reader,
		writer:  writer,
	}
}

// Counts all matches and prints matches for matchers af printing kind
func (m Matcher) PrintAndCount() error {
	s := bufio.NewScanner(m.reader)

	// Scan reader line by line
	for s.Scan() {
		line := s.Text()

		// Loop all matchers to count and maybe print line
		for i, match := range m.matches { //i := 0; i < len(m.matches); i++ {

			hit, err := regexp.Match(match.Pattern, []byte(line))
			if err != nil {
				return err
			}

			if hit {
				match.Count++

				// Print if that kind of matcher
				if match.Kind == PrintMatch || match.Kind == PrintCountMatch {
					fmt.Printf("<%d> %v\n", i, line)
				}
			}
		}
	}

	return nil
}

// Print the counts of all counting matchers
func (m Matcher) PrintCounts() {

	for i, match := range m.matches {
		if match.Kind == CountMatch || match.Kind == PrintCountMatch {
			fmt.Fprintf(m.writer, "Match %q <%d> got %d matches\n", match.Pattern, i, match.Count)
		}
	}

}

// Converts a cammand line flag to matcher kind
func FlagToMatchKind(flag string) (k MatchKind, err error) {

	switch flag {
	case "-p":
		return PrintMatch, nil
	case "-c":
		return CountMatch, nil
	case "-pc":
		return PrintCountMatch, nil
	}

	return k, errors.New("Unknown flag: " + flag)
}
