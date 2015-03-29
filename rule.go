package css

type CSSRule struct {
	property  string
	value     string
	immutable bool
	defLine   int
	before    string
	after     string
}

func NewRule(property []byte, line int, immutable bool) *CSSRule {
	before, prop, after := parseBytes(property)
	return &CSSRule{
		property:  string(prop),
		defLine:   line,
		immutable: immutable,
		before:    string(before),
		after:     string(after),
	}
}

func (r *CSSRule) IsSpecialProperty() (special bool) {
	// Case IE's filer
	if r.property == "filter" {
		special = true
	}

	return
}

func (r *CSSRule) SetValue(value []byte) {
	r.value = string(value)
}
