# trieregex

`trieregex` builds compact regular expressions from code points stored within a trie.

### Usage

```go
package main

import (
  "log"
	
  "github.com/n3integration/trieregex"
)

func main() {
  trie := trieregex.NewTrie()
  trie.Add("ham")
  trie.Add("spam")
  trie.Add("spearfish")

  re, err := trie.ToRegex()
  if err != nil {
    // handle err
  }

  if re.String() == `^(?:ham|sp(?:am|earfish))$` {
    log.Println(`🎉`)
  }
}
```

##### License

```text
   Copyright 2022 n3integration

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```
