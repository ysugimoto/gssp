package css

type CSSValue struct {
	Value     string
	DefLine   int
	Point     int
	Before    string
	After     string
	Semicolon bool
	RawData   string
	//RawData   []byte
}

func NewValue(val []byte, line, point int, semicolon bool) *CSSValue {
	before, value, after := parseBytes(val)
	return &CSSValue{
		Value:     string(value),
		DefLine:   line,
		Point:     point,
		Before:    string(before),
		After:     string(after),
		RawData:   string(val),
		Semicolon: semicolon,
	}
}
