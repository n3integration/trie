package trie

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// ToRegex converts the underlying trie structure into a bounded regular expression
func (t *Trie[T]) ToRegex() (*regexp.Regexp, error) {
	if len(t.nodes) == 0 {
		return nil, errors.New("trie is empty")
	}

	rePattern := t.Pattern()
	return regexp.Compile(fmt.Sprintf("^%s$", rePattern))
}

// Pattern exposes the trie's regular expression fragment as a string
func (t *Trie[T]) Pattern() string {
	if len(t.nodes) == 0 {
		return ""
	}
	if len(t.nodes) == 1 {
		key := t.keys()[0]
		return t.nodes[key].makeSequence(key)
	}

	sequences := make([]string, 0)
	for key, trie := range t.nodes {
		sequences = append(sequences, trie.makeSequence(key))
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

func (t *Trie[T]) makeSequence(key rune) string {
	return fmt.Sprintf("%s%s", regexp.QuoteMeta(string(key)), t.Pattern())
}
