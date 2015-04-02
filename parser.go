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

	crlf      = regexp.MustCompile("\r\n")
	trimRegex = regexp.MustCompile("([\n\t ]*)([a-zA-Z0-9-]+)([;\n ]*)")
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

	for _, css := range c.definitions {
		pp.Print(css)
	}

	return ""
}

func (c *CSSParser) processEscapeSequence() {
	c.skipping = true
	c.isEscaping = true
	c.stack = append(c.stack, ESCAPE_SEQUENCE)
}

func (c *CSSParser) processLineFeed(index, point *int) {
	val := bytes.Trim(c.stack, ";:\n\t ")
	if len(val) > 0 {
		if !c.inSelector && val[0] == CONTROL_SIGNATURE {
			c.globalDefinitions = append(c.globalDefinitions, NewDefinition(
				NewSelector(c.stack),
				*index,
				*point-len(c.stack)+1,
			))
			c.stack = []byte{}
		} else if defRule != nil {
			defRule.SetValue(c.stack, *index, *point-len(c.stack)+1)
			defList.GetLastChild().AddRule(defRule)
			defRule = nil
			c.stack = []byte{}
		}
	}
}

func (c *CSSParser) processSelectorOpen(index, point *int) {
	def := NewDefinition(
		NewSelector(c.stack),
		*index,
		*point-len(c.stack)+1,
	)
	defList.AddDefinition(def)
	c.inSelector = true
	c.stack = []byte{}
}

func (c *CSSParser) processPropertySeparator(index, point *int) {
	if defRule != nil && defRule.IsSpecialProperty() || !c.inSelector {
		c.stack = append(c.stack, PROPERTY_SEPARSTOR)
		return
	}
	defRule = NewRule(
		c.stack,
		*index,
		*point-len(c.stack)+1,
	)
	c.stack = []byte{}
}

func (c *CSSParser) processValueEnd(index, point *int) {
	if !c.inSelector {
		c.globalDefinitions = append(c.globalDefinitions, NewDefinition(
			NewSelector(c.stack),
			*index,
			*point-len(c.stack)+1,
		))
		c.stack = []byte{}
		return
	}
	defRule.SetValue(c.stack, *index, *point-len(c.stack)+1)
	defList.GetLastChild().AddRule(defRule)
	defRule = nil
	c.stack = []byte{}
}

func (c *CSSParser) processSelectorClose(index, point *int) {
	cDef := defList.GetLastChild()
	if defRule != nil {
		defRule.SetValue(c.stack, *index, *point-len(c.stack)+1)
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
			c.processEscapeSequence()
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
			c.processLineFeed(&index, &point)
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
			c.processSelectorOpen(&index, &point)
			continue
		case PROPERTY_SEPARSTOR:
			if c.quoting {
				c.stack = append(c.stack, line[point])
				continue
			}
			c.processPropertySeparator(&index, &point)
			continue
		case VALUE_END:
			if c.quoting || c.skipping {
				c.stack = append(c.stack, line[point])
				continue
			}
			c.processValueEnd(&index, &point)
			continue
		case SELECTOR_CLOSE:
			if c.quoting {
				c.stack = append(c.stack, line[point])
				continue
			}
			c.processSelectorClose(&index, &point)
			continue
		}

		if c.isEscaping {
			c.isEscaping = false
			c.skipping = false
		}
		c.stack = append(c.stack, line[point])
	}

	// check remains
	if defRule != nil {
		defRule.SetValue(c.stack, index, len(line)-len(c.stack)+1)
		defList.GetLastChild().AddRule(defRule)
		defRule = nil
	}
	if defList.Remains() {
		c.definitions = append(c.definitions, defList.GetLastChild())
	}
}
