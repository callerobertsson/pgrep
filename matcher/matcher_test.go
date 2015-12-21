package matcher

import (
	"errors"
	"testing"
)

func TestFlagToMatch(t *testing.T) {

	cases := []struct {
		flag         string
		expectedErr  error
		expectedKind MatchKind
	}{
		{"-p", nil, PrintMatch},
		{"-c", nil, CountMatch},
		{"-pc", nil, PrintCountMatch},
		{"qk", errors.New("Unknown flag: qk"), -1},
	}

	for _, c := range cases {
		kind, err := FlagToMatchKind(c.flag)

		// If expecting error
		if c.expectedErr != nil {
			if err == nil {
				t.Errorf("FlagToMatchKind(%q) succeded, expected error: %q\n", c.flag, c.expectedErr.Error())
			} else if c.expectedErr.Error() != err.Error() {
				t.Errorf("FlagToMatchKind(%q) unexpected error: %v", c.flag, err.Error())
			}
			continue
		}

		// Check result
		if kind != c.expectedKind {
			t.Errorf("FlagToMatchKind(%q) == %v, expected %v\n",
				c.flag, kind, c.expectedKind)
		}
	}
}
