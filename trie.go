package trie

import (
	"reflect"
)

// Trie provides a basic retrieval structure that can hold any valid series
// of runes
type Trie[T comparable] struct {
	nodes map[rune]*Trie[T]
	value T
}

// New initializes an empty trie data structure
func New[T comparable]() *Trie[T] {
	return &Trie[T]{}
}

// NewSet returns a trie set
func NewSet() *Trie[struct{}] {
	return &Trie[struct{}]{}
}

// Add stores the key/value within the trie; returns true if a previous value
// is overwritten; false, otherwise
func (t *Trie[T]) Add(key string, value T) bool {
	if len(key) == 0 {
		return false
	}

	current := t
	for _, r := range key {
		node, ok := current.nodes[r]
		if !ok {
			if current.nodes == nil {
				current.nodes = map[rune]*Trie[T]{}
			}
			node = new(Trie[T])
			current.nodes[r] = node
		}
		current = node
	}

	previous := current.value
	current.value = value
	return previous != value
}

// Get the value associated with the provided key; otherwise, return the empty
// value if not found
func (t *Trie[T]) Get(key string) (val T, vOK bool) {
	current := t
	for _, r := range key {
		node, ok := current.nodes[r]
		if !ok {
			return val, false
		}
		current = node
	}
	return current.value, true
}

// Len returns the total number of nodes
func (t *Trie[T]) Len() int {
	if len(t.nodes) == 0 || !reflect.ValueOf(t.value).IsZero() {
		return 1
	}

	count := 0
	for _, trie := range t.nodes {
		count += trie.Len()
	}
	return count
}

func (t *Trie[T]) keys() []rune {
	var keys []rune
	for key := range t.nodes {
		keys = append(keys, key)
	}
	return keys
}
