package sidebar

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var (
	h2Re = regexp.MustCompile(`^#{2}\s+(.*)`)
	h3Re = regexp.MustCompile(`^#{3}\s+(.+)`)

	h2HtmlRe = regexp.MustCompile(`^<h2[^>]*>(.*)</h2>`)
	h3HtmlRe = regexp.MustCompile(`^<h3[^>]*>(.*)</h3>`)
)

type Item struct {
	Name     string  `json:"name"`
	Children []Child `json:"children"`
}

type Child struct {
	Name string `json:"name"`
}

type Result struct {
	sidebar  []Item
	lastIsH2 bool
}

func AddIdForH2H3(html string) string {
	var newContent []string
	reader := strings.NewReader(html)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		addID(&newContent, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strings.Join(newContent, "\n")
}

func GetSidebar(strMkd string) []Item2 {
	var result Result
	parse(&result, strMkd)
	return fixItem(result.sidebar)
}

func parse(result *Result, mkdStr string) {
	reader := strings.NewReader(mkdStr)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		scan(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func scan(result *Result, text string) {
	if len(text) == 0 {
		return
	} else if text[0] != '#' {
		return
	}

	//fmt.Println("text:", text)
	arr := h2Re.FindAllStringSubmatch(text, -1)
	if len(arr) > 0 {
		result.lastIsH2 = true
		h2Txt := arr[0][1]
		//fmt.Println("h2:", h2Txt)
		item := Item{Name: h2Txt}
		result.sidebar = append(result.sidebar, item)
		return
	}

	arr1 := h3Re.FindAllStringSubmatch(text, -1)
	if len(arr1) > 0 {
		h3Txt := arr1[0][1]
		//fmt.Println("h3:", h3Txt)
		if !result.lastIsH2 {
			item := Item{Name: h3Txt}
			result.sidebar = append(result.sidebar, item)
		} else {
			child := Child{Name: h3Txt}
			lastItem := result.sidebar[len(result.sidebar)-1]
			lastItem.Children = append(lastItem.Children, child)
			result.sidebar[len(result.sidebar)-1] = lastItem
		}
	}
}

func addID(content *[]string, line string) {
	//fmt.Println(line)
	var matched1, matched2 bool

	arr := h2HtmlRe.FindAllStringSubmatch(line, -1)
	if len(arr) > 0 {
		h2Txt := arr[0][1]
		//fmt.Println("h2:", fmt.Sprintf(`%s id='%s'%s`, line[0:3], h2Txt, line[3:]))
		*content = append(*content, fmt.Sprintf(`%s id='%s'%s`, line[0:3], h2Txt, line[3:]))
		return
	} else {
		matched1 = false
	}

	arr1 := h3HtmlRe.FindAllStringSubmatch(line, -1)
	if len(arr1) > 0 {
		h3Txt := arr1[0][1]
		//fmt.Println("h3:", fmt.Sprintf(`%s id='%s'%s`, line[0:3], h3Txt, line[3:]))
		*content = append(*content, fmt.Sprintf(`%s id='%s'%s`, line[0:3], h3Txt, line[3:]))
		return
	} else {
		matched2 = false
	}

	if !matched1 && !matched2 {
		*content = append(*content, line)
	}
}
