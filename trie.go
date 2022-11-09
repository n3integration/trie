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

// NewTrie initializes an empty trie data structure
func NewTrie() *Trie {
	return &Trie{}
}

// Add stores a string within the underlying data structure
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

	rePattern := t.Pattern()
	return regexp.Compile(fmt.Sprintf("^%s$", rePattern))
}

// Pattern exposes the trie's regular expression fragment as a string
func (t *Trie) Pattern() string {
	if len(t.nodes) == 0 {
		return ""
	}
	if len(t.nodes) == 1 {
		key := t.keys()[0]
		return fmt.Sprintf("%s%s", regexp.QuoteMeta(string(key)), t.nodes[key].Pattern())
	}

	sequences := make([]string, 0)
	for key, trie := range t.nodes {
		sequences = append(sequences, fmt.Sprintf("%s%s", regexp.QuoteMeta(string(key)), trie.Pattern()))
	}

	if len(sequences) == 1 {
		return sequences[0]
	}
	subsequence := strings.Join(sequences, "")
	if len(sequences) == len(subsequence) {
		return fmt.Sprintf("[%s]", subsequence)
	}
	sort.Slice(sequences, func(i, j int) bool {
		diff := strings.Compare(sequences[i], sequences[j])
		if diff == 0 {
			return len(sequences[i]) < len(sequences[j])
		}
		return diff < 0
	})
	return fmt.Sprintf("(?:%s)", strings.Join(sequences, "|"))
}

func (t *Trie) keys() []rune {
	var keys []rune
	for key := range t.nodes {
		keys = append(keys, key)
	}
	return keys
}
