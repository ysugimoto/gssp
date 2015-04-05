package css

type CSSSelector struct {
	Selector        string
	ControlSelector bool
	Before          string
	After           string
	//RawData         []byte
	RawData string
}

func NewSelector(selBytes []byte) *CSSSelector {
	before, selector, after := parseBytes(selBytes)
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
	}
}

func (s *CSSSelector) String() string {
	return s.Selector
}

func (s *CSSSelector) IsControlSelector() bool {
	return s.ControlSelector
}
