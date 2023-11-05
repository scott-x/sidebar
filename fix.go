package sidebar

import (
	"github.com/russross/blackfriday/v2"
	"strings"
)

type Item2 struct {
	ID       string   `json:"id"`
	Value    string   `json:"value"`
	Children []Child2 `json:"children"`
}

type Child2 struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func fixItem(items []Item) []Item2 {
	var item2s []Item2
	for _, item := range items {
		plain := item.Name
		bs := blackfriday.Run([]byte(plain))
		str := trimP(string(bs))
		id := fixID(str)
		//a := fmt.Sprintf(`<a href='#%s'>%s</a>`, id, str)
		item2 := Item2{
			ID:    id,
			Value: str,
		}
		children := item.Children
		if len(children) > 0 {
			for _, child := range children {
				plain := child.Name
				bs := blackfriday.Run([]byte(plain))
				str := trimP(string(bs))
				id := fixID(str)
				//a := fmt.Sprintf(`<a href='#%s'>%s</a>`, id, str)
				child2 := Child2{
					ID:    id,
					Value: str,
				}
				item2.Children = append(item2.Children, child2)
			}
		}
		item2s = append(item2s, item2)
	}
	return item2s
}

// eg: 11 <code>&lt;Input/&gt;</code> &amp; <code>vue</code>
func fixID(htmStr string) string {
	id := strings.ReplaceAll(htmStr, "&lt;", `<`)
	id = strings.ReplaceAll(id, "&gt;", `>`)
	return strings.ReplaceAll(id, "&amp;", "&")
}

// trim <p>...</p>
func trimP(htmStr string) string {
	//remove <p></p>1
	str := strings.TrimPrefix(strings.TrimSpace(string(htmStr)), `<p>`)
	return strings.TrimSuffix(str, `</p>`)
}
