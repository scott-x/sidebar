# sidebar

parse sidebar for markdown

### struct & API

```go
type Item2 struct {
	ID       string   `json:"id"`
	Value    string   `json:"value"`
	Children []Child2 `json:"children"`
}
```
- `func GetSidebar(strMkd string) []Item2`:  get sidebar structure.
- `func AddIdForH2H3(html string) string`: normally markdown pkg won't add id for h2 & h3, this function can do it for you.

### demo

```go
//GetSidebar: parse markdown text
package main

import (
	"fmt"
	"github.com/scott-x/sidebar"
	"io/ioutil"
)

func main() {
	bs, _ := ioutil.ReadFile("./test.md")
	items := sidebar.GetSidebar(string(bs))
	for _, item := range items {
		fmt.Println(item)
	}
}
```

```go
//AddIdForH2H3: : parse html text
package main

import (
	"fmt"
	"github.com/scott-x/sidebar"
	"io/ioutil"
)

func main() {
	bs, _ := ioutil.ReadFile("./test.txt")
	res := sidebar.AddIdForH2H3(string(bs))
	fmt.Println(res)
}
```