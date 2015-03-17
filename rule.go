package css

type CSSRule struct {
	property  string
	value     string
	immutable bool
	defLine   int
}

func NewRule(property []byte, line int, immutable bool) *CSSRule {
	return &CSSRule{
		property:  string(property),
		defLine:   line,
		immutable: immutable,
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
