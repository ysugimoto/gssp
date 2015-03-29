package css

import (
	"bytes"
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

	PARENTHEIS_LEFT  = 40
	PARENTHEIS_RIGHT = 41

	LINE_FEED = 10

	ESCAPE_SEQUENCE = 92
)

var (
	defList = NewDefinitionList()
	defRule *CSSRule

	crlf = regexp.MustCompile("\r\n")
)

type CSSParser struct {
	lineNumber        int
	definitions       []*CSSDefinition
	comment           bool
	quoting           bool
	singleQuote       bool
	doubleQuote       bool
	inSelector        bool
	skipping          bool
	isEscaping        bool
	globalDefinitions []*CSSDefinition
	stack             []byte
}

func NewParser() *CSSParser {
	return &CSSParser{
		lineNumber:  0,
		definitions: make([]*CSSDefinition, 0),
		stack:       []byte{},
	}
}

func (c *CSSParser) Parse(buffer []byte) string {
	LF := []byte("\n")
	buffer = crlf.ReplaceAll(buffer, LF)

	c.execParse(buffer)

	pp.Print(c)

	return ""
}

func (c *CSSParser) execParse(line []byte) {
	index := 1
	for point := 0; point < len(line); point++ {
		if c.isCommentStart(line, point) {
			c.comment = true
			continue
		}
		if c.isCommentEnd(line, point) {
			c.comment = false
			point++
			continue
		}

		if c.comment {
			continue
		}

		switch line[point] {
		case ESCAPE_SEQUENCE:
			c.skipping = true
			c.isEscaping = true
			c.stack = append(c.stack, ESCAPE_SEQUENCE)
			continue
		case PARENTHEIS_LEFT:
			if !c.quoting {
				c.skipping = true
			}
		case PARENTHEIS_RIGHT:
			if !c.quoting {
				c.skipping = false
			}
		case LINE_FEED:
			val := bytes.Trim(c.stack, ";:\n\t ")
			c.stack = []byte{}
			if len(val) > 0 {
				if !c.inSelector && val[0] == CONTROL_SIGNATURE {
					c.globalDefinitions = append(c.globalDefinitions, NewDefinition(
						NewSelector(val),
						index+1,
					))
				} else if defRule != nil {
					defRule.SetValue(val)
					defList.GetLastChild().AddRule(defRule)
					defRule = nil
				}
			}
			index++
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
			if c.quoting {
				c.stack = append(c.stack, line[point])
				break
			}
			sel := bytes.TrimSpace(c.stack)
			c.stack = []byte{}
			def := NewDefinition(
				NewSelector(sel),
				index+1,
			)
			defList.AddDefinition(def)
			c.inSelector = true
			continue
		case PROPERTY_SEPARSTOR:
			if c.quoting {
				c.stack = append(c.stack, line[point])
				continue
			}
			if defRule != nil && defRule.IsSpecialProperty() || !c.inSelector {
				c.stack = append(c.stack, PROPERTY_SEPARSTOR)
				continue
			}
			defRule = NewRule(
				bytes.Trim(c.stack, ";:\n\t "),
				index+1,
				false,
			)
			c.stack = []byte{}
		case VALUE_END:
			if c.quoting || c.skipping {
				c.stack = append(c.stack, line[point])
				continue
			}
			val := bytes.Trim(c.stack, ";:\n\t ")
			c.stack = []byte{}
			if len(val) == 0 {
				continue
			}
			if !c.inSelector {
				c.globalDefinitions = append(c.globalDefinitions, NewDefinition(
					NewSelector(val),
					index+1,
				))
				point++
				continue
			}
			defRule.SetValue(val)
			defList.GetLastChild().AddRule(defRule)
			defRule = nil
		case SELECTOR_CLOSE:
			if c.quoting {
				c.stack = append(c.stack, line[point])
				continue
			}
			cDef := defList.GetLastChild()
			if defRule != nil {
				defRule.SetValue(bytes.Trim(c.stack, ";:\n\t "))
				cDef.AddRule(defRule)
				defRule = nil
				c.stack = []byte{}
			}
			defList.Remove()
			if defList.Remains() {
				defList.GetLastChild().AddControl(cDef)
			} else {
				c.definitions = append(c.definitions, cDef)
			}
			c.inSelector = false
		}

		if c.isEscaping {
			c.isEscaping = false
			c.skipping = false
		}
		c.stack = append(c.stack, line[point])
	}

	if defRule != nil {
		defRule.SetValue(bytes.Trim(c.stack, ";: "))
		defList.GetLastChild().AddRule(defRule)
		defRule = nil
	}
	if defList.Remains() {
		c.definitions = append(c.definitions, defList.GetLastChild())
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
