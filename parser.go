package gssp

import (
	"bytes"
)

// Byte number of signatures
const (
	SELECTOR_OPEN      = 123 // "{"
	SELECTOR_CLOSE     = 125 // "}"
	PROPERTY_SEPARSTOR = 58  // ":"
	VALUE_END          = 59  // ";"
	COMMENT_SLASH      = 47  // "/"
	COMMENT_STAR       = 42  // "*"
	CONTROL_SIGNATURE  = 64  // "@"
	DOUBLE_QUOTE       = 34  // "\""
	SINGLE_QUOTE       = 39  // "'"
	PARENTHEIS_LEFT    = 40  // "("
	PARENTHEIS_RIGHT   = 41  // ")"
	LINE_FEED          = 10  // "\n"
	CARRIAGE_RETURN    = 13  // "\r"
	ESCAPE_SEQUENCE    = 92  // "\"
)

type CSSParser struct {
	definitions  *CSSDefinitionList
	comment      bool
	quoting      bool
	singleQuote  bool
	doubleQuote  bool
	inSelector   bool
	skipping     bool
	isEscaping   bool
	inParentheis bool
	defTree      *CSSDefinitionTree
	defRule      *CSSRule
	charPoint    int
	stack        []byte
}

func NewParser() *CSSParser {
	return &CSSParser{
		definitions: NewDefinitionList(),
		stack:       []byte{},
		defTree:     NewDefinitionTree(),
		defRule:     nil,
		charPoint:   0,
	}
}

func (c *CSSParser) Parse(buffer []byte) CSSParseResult {
	c.execParse(buffer)

	return CSSParseResult{
		data: c.definitions.Get(),
	}
}

func (c *CSSParser) processEscapeSequence() {
	if !c.skipping {
		c.skipping = true
		c.isEscaping = true
	} else {
		c.skipping = false
		c.isEscaping = false
	}

	c.stack = append(c.stack, ESCAPE_SEQUENCE)
}

func (c *CSSParser) processLineFeed(index *int) {
	val := bytes.Trim(c.stack, ";:\n\t ")
	if len(val) > 0 {
		if !c.inSelector && val[0] == CONTROL_SIGNATURE {
			def := NewDefinition(
				NewSelector(c.stack),
				*index,
				c.charPoint,
			)
			c.definitions.Add(def)
			c.stack = []byte{}
		} else if c.defRule != nil {
			c.defRule.SetValue(
				c.stack,
				*index,
				c.charPoint,
				false,
			)
			c.defTree.GetLastChild().AddRule(c.defRule)
			c.defRule = nil
			c.stack = []byte{}
		}
	}
	c.charPoint = 0
}

func (c *CSSParser) processSelectorOpen(index *int) {
	def := NewDefinition(
		NewSelector(c.stack),
		*index,
		c.charPoint,
	)
	c.defTree.AddDefinition(def)
	c.inSelector = true
	c.stack = []byte{}
}

func (c *CSSParser) processPropertySeparator(index *int) {
	if c.defRule != nil && c.defRule.IsSpecialProperty() || !c.inSelector {
		c.stack = append(c.stack, PROPERTY_SEPARSTOR)
		return
	}
	c.defRule = NewRule(
		c.stack,
		*index,
		c.charPoint,
	)

	c.stack = []byte{}
}

func (c *CSSParser) processValueEnd(index *int) {
	if !c.inSelector {
		def := NewDefinition(
			NewSelector(c.stack),
			*index,
			c.charPoint,
		)
		c.definitions.Add(def)
		c.stack = []byte{}
		return
	}
	if !isEmptyStack(c.stack) {
		c.defRule.SetValue(
			c.stack,
			*index,
			c.charPoint,
			true,
		)
		c.defTree.GetLastChild().AddRule(c.defRule)
		c.defRule = nil
	}
	c.stack = []byte{}
}

func (c *CSSParser) processSelectorClose(index *int) {
	cDef := c.defTree.GetLastChild()
	if c.defRule != nil {
		c.defRule.SetValue(
			c.stack,
			*index,
			c.charPoint,
			false,
		)
		cDef.AddRule(c.defRule)
		c.defRule = nil
	}
	c.defTree.Remove()
	if c.defTree.Remains() {
		c.defTree.GetLastChild().AddControl(cDef)
	} else {
		c.definitions.Add(cDef)
	}
	c.inSelector = false
	c.stack = []byte{}
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
		c.charPoint++
		if c.isCommentStart(line, point) {
			c.comment = true
			c.stack = append(c.stack, line[point])
			continue
		}
		if c.isCommentEnd(line, point) {
			c.comment = false
			c.stack = append(c.stack, COMMENT_STAR, COMMENT_SLASH)
			point++
			c.charPoint++
			continue
		}

		if c.comment {
			c.stack = append(c.stack, line[point])
			continue
		}

		switch line[point] {
		case ESCAPE_SEQUENCE:
			c.processEscapeSequence()
			continue
		case PARENTHEIS_LEFT:
			if !c.quoting {
				c.inParentheis = true
			}
		case PARENTHEIS_RIGHT:
			if !c.quoting {
				c.inParentheis = false
			}
		case LINE_FEED:
			c.processLineFeed(&index)
			index++
		case DOUBLE_QUOTE:
			if c.skipping || c.singleQuote {
				break
			}
			if c.doubleQuote {
				c.quoting = false
				c.doubleQuote = false
			} else {
				c.quoting = true
				c.doubleQuote = true
			}
		case SINGLE_QUOTE:
			if c.skipping || c.doubleQuote {
				break
			}
			if c.singleQuote {
				c.quoting = false
				c.singleQuote = false
			} else {
				c.quoting = true
				c.singleQuote = true
			}
		case SELECTOR_OPEN:
			if c.quoting || c.skipping || c.inParentheis {
				break
			}
			c.processSelectorOpen(&index)
			continue
		case PROPERTY_SEPARSTOR:
			if c.quoting || c.skipping || c.inParentheis {
				break
			}
			c.processPropertySeparator(&index)
			continue
		case VALUE_END:
			if c.quoting || c.skipping || c.inParentheis {
				break
			}
			c.processValueEnd(&index)
			continue
		case SELECTOR_CLOSE:
			if c.quoting || c.skipping || c.inParentheis {
				break
			}
			c.processSelectorClose(&index)
			continue
		}

		if c.isEscaping {
			c.isEscaping = false
			c.skipping = false
		}
		c.stack = append(c.stack, line[point])
	}

	// check remains
	if c.defRule != nil {
		c.defRule.SetValue(
			c.stack,
			index,
			c.charPoint,
			false,
		)
		c.defTree.GetLastChild().AddRule(c.defRule)
		c.defRule = nil
	}
	if c.defTree.Remains() {
		c.definitions.Add(c.defTree.GetLastChild())
	}
}
