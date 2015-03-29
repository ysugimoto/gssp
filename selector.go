package css

type CSSSelector struct {
	selector        string
	controlSelector bool
}

func NewSelector(selBytes []byte) *CSSSelector {
	var isControl bool

	if len(selBytes) > 0 && selBytes[0] == CONTROL_SIGNATURE {
		isControl = true
	} else {
		isControl = false
	}
	return &CSSSelector{
		selector:        string(selBytes),
		controlSelector: isControl,
	}
}

func (s *CSSSelector) String() string {
	return s.selector
}

func (s *CSSSelector) IsControlSelector() bool {
	return s.controlSelector
}
