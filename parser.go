package css

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/k0kubun/pp"
)

// Byte number of signatures
const (
	SELECTOR_OPEN      = 123
	SELECTOR_CLOSE     = 125
	PROPERTY_SEPARSTOR = 58
	VALUE_END          = 59

	COMMENT_SLASH = 47
	COMMENT_STAR  = 42

	CONTROL_SIGNATURE = 64

	DOUBLE_QUOTE = 34
	SINGLE_QUOTE = 39

	LINE_FEED = 10
)

var (
	defList = NewDefinitionList()
	defRule *CSSRule

	crlf = regexp.MustCompile("\r\n")
)

type CSSParser struct {
	lineNumber        int
	charNumber        int
	definitions       []*CSSDefinition
	comment           bool
	quoting           bool
	singleQuote       bool
	doubleQuote       bool
	inSelector        bool
	globalDefinitions []*CSSDefinition
}

func NewParser() *CSSParser {
	return &CSSParser{
		lineNumber:  0,
		charNumber:  0,
		definitions: make([]*CSSDefinition, 0),
	}
}

func (c *CSSParser) Parse(buffer []byte) string {
	LF := []byte("\n")
	buffer = crlf.ReplaceAll(buffer, LF)

	c.execParse(buffer)
	//lines := bytes.Split(buffer, []byte("\n"))

	//for index, line := range lines {
	//	c.lineNumber = index
	//	c.charNumber = 0
	//	c.parseLine(line, index)
	//}

	//if len(c.defTree) > 0 {
	//	pp.Print(c.currentDef)
	//	fmt.Println("Syntax Error: Unexpected end of css")
	//	return ""
	//} else if c.currentRule != nil {
	//	pp.Print(c.currentRule)
	//	fmt.Println("Syntax Error: Unexpected end of css")
	//	return ""
	//}

	//pp.Print(c)

	return ""
}

func (c *CSSParser) execParse(line []byte) {
	index := 1
	for point := 0; point < len(line); point++ {
		if c.isCommentStart(line, point) {
			c.comment = true
			fmt.Println("comment start")
			continue
		}
		if c.isCommentEnd(line, point) {
			c.comment = false
			c.charNumber = point + 2
			fmt.Println("comment end")
			continue
		}

		if line[point] == LINE_FEED {
			val := bytes.Trim(line[c.charNumber:point], ";:\n\t ")
			if len(val) > 0 {
				if !c.inSelector && val[0] == CONTROL_SIGNATURE {
					c.globalDefinitions = append(c.globalDefinitions, NewDefinition(
						NewSelector(val),
						index,
					))
				}
			}
			index++
			continue
		}

		if c.comment {
			c.charNumber++
			continue
		}

		switch line[point] {
		case DOUBLE_QUOTE:
			if c.doubleQuote && c.quoting {
				c.quoting = false
				c.doubleQuote = false
			} else {
				c.quoting = true
				c.doubleQuote = true
			}
		case SINGLE_QUOTE:
			if c.singleQuote && c.quoting {
				c.quoting = false
				c.singleQuote = false
			} else {
				c.quoting = true
				c.singleQuote = true
			}
		case SELECTOR_OPEN:
			fmt.Println("open")
			if c.quoting {
				fmt.Println("open quoting")
				continue
			}
			sel := bytes.TrimSpace(line[c.charNumber:point])
			fmt.Println(c.charNumber)
			fmt.Println(point)
			if len(sel) == 0 {
				fmt.Println("empty")
				c.charNumber = point
				continue
			}
			pp.Print(sel)
			def := NewDefinition(
				NewSelector(sel),
				index,
			)
			if defList.HasParent() {
				defList.GetLastChild().AddControl(def)
			} else {
				defList.AddDefinition(def)
			}
			c.inSelector = true
			c.charNumber = point + 1
		case PROPERTY_SEPARSTOR:
			if c.quoting {
				continue
			}
			if defRule != nil && defRule.IsSpecialProperty() || !c.inSelector {
				continue
			}
			defRule = NewRule(
				bytes.Trim(line[c.charNumber:point], ";:\n\t "),
				index,
				false,
			)
			c.charNumber = point + 1
		case VALUE_END:
			if c.quoting {
				continue
			}
			val := bytes.Trim(line[c.charNumber:point], ";:\n\t ")
			if len(val) == 0 {
				continue
			}
			if !c.inSelector {
				c.globalDefinitions = append(c.globalDefinitions, NewDefinition(
					NewSelector(val),
					index,
				))
				point++
				c.charNumber = point
				continue
			}
			fmt.Println(string(val))
			defRule.SetValue(val)
			defList.GetLastChild().AddRule(defRule)
			defRule = nil
			c.charNumber = point + 1
		case SELECTOR_CLOSE:
			fmt.Println("close")
			if c.quoting {
				continue
			}
			cDef := defList.GetLastChild()
			if defRule != nil {
				defRule.SetValue(bytes.Trim(line[c.charNumber:point], ";:\n\t "))
				cDef.AddRule(defRule)
				defRule = nil
			}
			defList.Remove()
			if defList.Remains() {
				defList.GetLastChild().AddControl(cDef)
			} else {
				c.definitions = append(c.definitions, cDef)
			}
			c.inSelector = false
			c.charNumber = point + 1
		}

		if point == len(line)-1 {
			if defRule != nil {
				defRule.SetValue(bytes.Trim(line[c.charNumber:], ";: "))
				defList.GetLastChild().AddRule(defRule)
				defRule = nil
			}
			if defList.Remains() {
				c.definitions = append(c.definitions, defList.GetLastChild())
				//defList.Remove()
			}
		}
	}
}

func (c *CSSParser) isCommentStart(line []byte, point int) (start bool) {
	if len(line) <= point+1 || c.quoting {
		return
	}

	if line[point] == COMMENT_SLASH && line[point+1] == COMMENT_STAR {
		start = true
	}

	return
}

func (c *CSSParser) isCommentEnd(line []byte, point int) (end bool) {
	if len(line) <= point+1 || c.quoting {
		return
	}

	if line[point] == COMMENT_STAR && line[point+1] == COMMENT_SLASH {
		end = true
	}

	return
}
