package css

type CSSDefinition struct {
	Selector *CSSSelector
	Rules    []*CSSRule
	Controls []*CSSDefinition
	DefLine  int
	Point    int
	Parent   *CSSDefinition
}

func NewDefinition(selector *CSSSelector, line, point int) *CSSDefinition {
	return &CSSDefinition{
		Selector: selector,
		DefLine:  line,
		Point:    point,
		Rules:    make([]*CSSRule, 0),
		Controls: make([]*CSSDefinition, 0),
	}
}

func (d *CSSDefinition) AddRule(rule *CSSRule) {
	d.Rules = append(d.Rules, rule)
}

func (d *CSSDefinition) AddControl(control *CSSDefinition) {
	d.Controls = append(d.Controls, control)
}

func (d *CSSDefinition) AddChild(def *CSSDefinition) {
	def.Parent = d
}

func (d *CSSDefinition) GetParent() *CSSDefinition {
	return d.Parent
}

func (d *CSSDefinition) IsControl() bool {
	return d.Selector.IsControlSelector()
}
