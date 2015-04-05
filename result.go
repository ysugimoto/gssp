package gssp

import "encoding/json"

type CSSParseResult struct {
	data []*CSSDefinition
}

func (c CSSParseResult) ToJSON() []byte {
	ret, _ := json.Marshal(c.data)

	return ret
}

func (c CSSParseResult) ToJSONString() string {
	ret, _ := json.Marshal(c.data)

	return string(ret)
}

func (c CSSParseResult) ToPrettyJSON() []byte {
	ret, _ := json.MarshalIndent(c.data, "", "  ")

	return ret
}

func (c CSSParseResult) ToPrettyJSONString() string {
	ret, _ := json.MarshalIndent(c.data, "", "  ")

	return string(ret)
}

func (c CSSParseResult) GetData() []*CSSDefinition {
	return c.data
}
