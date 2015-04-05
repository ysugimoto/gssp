package css

type CSSSelector struct {
	Selector        string `json:"selector"`
	ControlSelector bool   `json:"atrule"`
	Before          string `json:"before"`
	After           string `json:"after"`
	RawData         string `json:"raw"`
	RawOffset       int    `json:"-"`
}

func NewSelector(selBytes []byte) *CSSSelector {
	before, selector, after, offset := parseBytes(selBytes)
	var isControl bool

	if len(selector) > 0 && selector[0] == CONTROL_SIGNATURE {
		isControl = true
	} else {
		isControl = false
	}

	return &CSSSelector{
		Before:          string(before),
		Selector:        string(selector),
		ControlSelector: isControl,
		After:           string(after),
		RawData:         string(selBytes),
		RawOffset:       offset,
	}
}

func (s *CSSSelector) String() string {
	return s.Selector
}

func (s *CSSSelector) IsControlSelector() bool {
	return s.ControlSelector
}
