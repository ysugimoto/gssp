package css

type CSSSelector struct {
	selector        string
	controlSelector bool
	before          string
	after           string
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
		before:          string(before),
		selector:        string(selector),
		controlSelector: isControl,
		after:           string(after),
	}
}

func (s *CSSSelector) String() string {
	return s.selector
}

func (s *CSSSelector) IsControlSelector() bool {
	return s.controlSelector
}
