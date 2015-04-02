package css

type CSSRule struct {
	property string
	value    *CSSValue
	defLine  int
	point    int
	before   string
	after    string
}

func NewRule(property []byte, line, point int) *CSSRule {
	before, prop, after := parseBytes(property)
	return &CSSRule{
		property: string(prop),
		defLine:  line,
		point:    point,
		before:   string(before),
		after:    string(after),
	}
}

func (r *CSSRule) IsSpecialProperty() (special bool) {
	// Case IE's filer
	if r.property == "filter" {
		special = true
	}

	return
}

func (r *CSSRule) SetValue(value []byte, index, point int) {
	r.value = NewValue(value, index, point, false)
}
