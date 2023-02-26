package trie

import (
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	values := []string{
		"alabama",
		"alaska",
		"arizona",
		"arkansas",
		"california",
		"colorado",
		"connecticut",
		"delaware",
		"florida",
		"georgia",
		"hawaii",
		"idaho",
		"illinois",
		"indiana",
		"iowa",
		"kansas",
		"kentucky",
		"louisiana",
		"maine",
		"maryland",
		"massachusetts",
		"michigan",
		"minnesota",
		"mississippi",
		"missouri",
		"montana",
		"nebraska",
		"nevada",
		"new hampshire",
		"new jersey",
		"new mexico",
		"new york",
		"north carolina",
		"north dakota",
		"ohio",
		"oklahoma",
		"oregon",
		"pennsylvania",
		"rhode island",
		"south carolina",
		"south dakota",
		"tennessee",
		"texas",
		"utah",
		"vermont",
		"virginia",
		"washington",
		"west virginia",
		"wisconsin",
		"wyoming",
	}

	trie := NewSet()
	for _, value := range values {
		trie.Add(value, struct{}{})
	}

	re, err := trie.ToRegex()
	if err != nil {
		t.Fatalf("failed to convert trie into regular expression: %s", err)
	}

	t.Run("Exists", func(t *testing.T) {
		checkExistence(t, re, true, values...)
	})

	t.Run("Does not exist", func(t *testing.T) {
		checkExistence(t, re, false, "block island", "providence")
	})
}

func checkExistence(t *testing.T, re *regexp.Regexp, exists bool, values ...string) {
	for _, value := range values {
		matches := re.MatchString(value)
		if exists && !matches {
			t.Errorf("pattern failed to match: %s", value)
		} else if !exists && matches {
			t.Errorf("pattern incorrectly resulted in a match: %s", value)
		}
	}
}
