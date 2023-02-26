package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	expectedCapital := "carson city"

	trie := New[string]()
	trie.Add("nevada", expectedCapital)
	trie.Add("new hampshire", "concord")
	trie.Add("new mexico", "albuquerque")
	trie.Add("new york", "albany")

	t.Run("Get", func(t *testing.T) {
		if v, ok := trie.Get("nevada"); ok {
			if v != expectedCapital {
				t.Errorf("incorrect value returned: expected %q, but got %q", expectedCapital, v)
			}
		}

		if _, ok := trie.Get("rhode island"); ok {
			t.Error("incorrect value returned: expected false, but got true")
		}
	})

	t.Run("Len", func(t *testing.T) {
		expected := 4
		actual := trie.Len()
		if expected != actual {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
}
