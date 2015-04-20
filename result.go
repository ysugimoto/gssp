package gssp

import "encoding/json"

type CSSParseResult struct {
	data     []*CSSDefinition
	Lines    int
	RawBytes []byte
	Files    []string
}

func NewParseResult(data []*CSSDefinition) *CSSParseResult {
	return &CSSParseResult{
		data:     data,
		Files:    []string{},
		RawBytes: []byte{},
	}
}

func (c *CSSParseResult) Merge(result *CSSParseResult) {
	c.data = append(c.data, result.data...)
	c.Lines += result.Lines
	c.RawBytes = append(c.RawBytes, result.RawBytes...)
	c.Files = append(c.Files, result.Files...)
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
