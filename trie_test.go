package trieregex

import (
	"testing"
)

func Test(t *testing.T) {
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

	trie := NewTrie()
	for _, value := range values {
		trie.Add(value)
	}

	re, err := trie.ToRegex()
	if err != nil {
		t.Fatalf("failed to convert trie into regular expression: %s", err)
	}

	for _, value := range values {
		if !re.MatchString(value) {
			t.Errorf("Pattern failed to match: %s", value)
		}
	}
}
