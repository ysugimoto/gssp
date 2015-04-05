package css

type CSSValue struct {
	Value     string `json:"data"`
	DefLine   int    `json:"line"`
	Point     int    `json:"column"`
	Before    string `json:"before"`
	After     string `json:"after"`
	Semicolon bool   `json:"semicolon"`
	RawData   string `json:"raw"`
}

func NewValue(val []byte, line, point int, semicolon bool) *CSSValue {
	before, value, after, offset := parseBytes(val)
	return &CSSValue{
		Value:     string(value),
		DefLine:   line,
		Point:     point - offset,
		Before:    string(before),
		After:     string(after),
		RawData:   string(val),
		Semicolon: semicolon,
	}
}
