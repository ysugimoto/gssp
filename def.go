package css

type CSSDefinition struct {
	selector *CSSSelector
	rules    []*CSSRule
	controls []*CSSDefinition
	defLine  int
	parent   *CSSDefinition
}

func NewDefinition(selector *CSSSelector, line int) *CSSDefinition {
	return &CSSDefinition{
		selector: selector,
		defLine:  line,
		rules:    make([]*CSSRule, 0),
		controls: make([]*CSSDefinition, 0),
	}
}

func (d *CSSDefinition) AddRule(rule *CSSRule) {
	d.rules = append(d.rules, rule)
}

func (d *CSSDefinition) AddControl(control *CSSDefinition) {
	d.controls = append(d.controls, control)
}

func (d *CSSDefinition) AddChild(def *CSSDefinition) {
	def.parent = d
}

func (d *CSSDefinition) GetParent() *CSSDefinition {
	return d.parent
}

func (d *CSSDefinition) IsControl() bool {
	return d.selector.IsControlSelector()
}
