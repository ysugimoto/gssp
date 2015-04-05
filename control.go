package gssp

type CSSControl struct {
	*CSSRule
	selector *CSSSelector
	rules    []*CSSRule
	defLine  int
}

func NewControl(selector *CSSSelector, line int) *CSSControl {
	return &CSSControl{
		selector: selector,
		defLine:  line,
		rules:    make([]*CSSRule, 0),
	}
}

func (d *CSSControl) AddRule(rule *CSSRule) {
	d.rules = append(d.rules, rule)
}
