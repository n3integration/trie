package trieregex

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Trie provides a basic retrieval structure that can hold any valid series of runes
type Trie struct {
	nodes map[rune]*Trie
}

// NewTrie initializes an empty trie structure
func NewTrie() *Trie {
	return &Trie{}
}

// Add saves
func (t *Trie) Add(text string) {
	current := t
	for _, r := range text {
		node, ok := current.nodes[r]
		if !ok {
			if current.nodes == nil {
				current.nodes = map[rune]*Trie{}
			}
			node = new(Trie)
			current.nodes[r] = node
		}
		current = node
	}
}

// ToRegex converts the underlying trie structure into a regular expression
func (t *Trie) ToRegex() (*regexp.Regexp, error) {
	if len(t.nodes) == 0 {
		return nil, errors.New("trie is empty")
	}

	regex := t.regex()
	return regexp.Compile(fmt.Sprintf("^%s$", regex))
}

func (t *Trie) regex() string {
	if len(t.nodes) == 0 {
		return ""
	}
	if len(t.nodes) == 1 {
		key := t.keys()[0]
		return fmt.Sprintf("%s%s", string(key), t.nodes[key].regex())
	}

	sequences := make([]string, 0)
	for key, trie := range t.nodes {
		sequences = append(sequences, fmt.Sprintf("%s%s", string(key), trie.regex()))
	}

	if len(sequences) == 1 {
		return sequences[0]
	}
	if len(sequences) == len(strings.Join(sequences, "")) {
		return fmt.Sprintf("[%s]", strings.Join(sequences, ""))
	}
	sort.Slice(sequences, func(i, j int) bool {
		return len(sequences[i]) < len(sequences[j])
	})
	return fmt.Sprintf("(?:%s)", strings.Join(sequences, "|"))
}

func (t *Trie) keys() []rune {
	keys := make([]rune, 0)
	for key := range t.nodes {
		keys = append(keys, key)
	}
	return keys
}
