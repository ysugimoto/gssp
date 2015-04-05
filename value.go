package css

type CSSValue struct {
	Value     string
	DefLine   int
	Point     int
	Before    string
	After     string
	Semicolon bool
	RawData   string
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
