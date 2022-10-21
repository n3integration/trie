# trieregex

### Usage

```go
package main

import "github.com/n3integration/trieregex"

func main() {
	trie := NewTrie()
	trie.Add("ham")
	trie.Add("spam")
	trie.Add("spearfish")

	re, err := trie.ToRegex()
	if err != nil {
		... // handle err
	}
	
	// re.String() = ^(?:ham|sp(?:am|earfish))$
  ...
}
```
