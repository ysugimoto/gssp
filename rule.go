package css

type CSSRule struct {
	Property  string
	Value     *CSSValue
	DefLine   int
	Point     int
	Before    string
	After     string
	RawData   string
	RawPoint  int
	RawOffset int
}

func NewRule(property []byte, line, point int) *CSSRule {
	before, prop, after, offset := parseBytes(property)
	return &CSSRule{
		Property:  string(prop),
		DefLine:   line,
		Point:     point - offset,
		Before:    string(before),
		After:     string(after),
		RawData:   string(property),
		RawPoint:  point,
		RawOffset: offset,
	}
}

func (r *CSSRule) IsSpecialProperty() (special bool) {
	// Case IE's filer
	if r.Property == "filter" {
		special = true
	}

	// TODO: Other case?

	return
}

func (r *CSSRule) SetValue(value []byte, index, point int, semicolon bool) {
	r.Value = NewValue(value, index, point, semicolon)
}
