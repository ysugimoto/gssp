package gssp

type CSSRule struct {
	Property  string    `json:"property"`
	Value     *CSSValue `json:"value"`
	DefLine   int       `json:"line"`
	Point     int       `json:"column"`
	Before    string    `json:"before"`
	After     string    `json:"after"`
	RawData   string    `json:"raw"`
	RawPoint  int       `json:"-"`
	RawOffset int       `json:"-"`
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
