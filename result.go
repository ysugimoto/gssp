package gssp

import "encoding/json"

type CSSParseResult struct {
	data     []*CSSDefinition
	lines    int
	rawBytes []byte
	file     string
}

func NewParseResult(data []*CSSDefinition, file string, raw []byte, lines int) *CSSParseResult {
	return &CSSParseResult{
		data:     data,
		lines:    lines,
		file:     file,
		rawBytes: raw,
	}
}

func (c *CSSParseResult) Merge(result *CSSParseResult) {
	c.data = append(c.data, result.data...)
	println(len(c.data))
}

func (c *CSSParseResult) ToJSON() []byte {
	ret, _ := json.Marshal(c.data)

	return ret
}

func (c *CSSParseResult) ToJSONString() string {
	ret, _ := json.Marshal(c.data)

	return string(ret)
}

func (c *CSSParseResult) ToPrettyJSON() []byte {
	ret, _ := json.MarshalIndent(c.data, "", "  ")

	return ret
}

func (c *CSSParseResult) ToPrettyJSONString() string {
	ret, _ := json.MarshalIndent(c.data, "", "  ")

	return string(ret)
}

func (c *CSSParseResult) Getdata() []*CSSDefinition {
	return c.data
}
