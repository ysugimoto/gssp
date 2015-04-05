package css

type CSSRule struct {
	Property string
	Value    *CSSValue
	DefLine  int
	Point    int
	Before   string
	After    string
	RawData  string
	//RawData  []byte
}

func NewRule(property []byte, line, point int) *CSSRule {
	before, prop, after := parseBytes(property)
	return &CSSRule{
		Property: string(prop),
		DefLine:  line,
		Point:    point,
		Before:   string(before),
		After:    string(after),
		RawData:  string(property),
	}
}

func (r *CSSRule) IsSpecialProperty() (special bool) {
	// Case IE's filer
	if r.Property == "filter" {
		special = true
	}

	return
}

func (r *CSSRule) SetValue(value []byte, index, point int, semicolon bool) {
	r.Value = NewValue(value, index, point, semicolon)
}
