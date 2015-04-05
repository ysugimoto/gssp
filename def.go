package css

type CSSDefinition struct {
	Selector *CSSSelector     `json:"selector"`
	Rules    []*CSSRule       `json:"rules"`
	Controls []*CSSDefinition `json:"controls"`
	DefLine  int              `json:"line"`
	Point    int              `json:"column"`
	Parent   *CSSDefinition   `json:"-"`
}

func NewDefinition(selector *CSSSelector, line, point int) *CSSDefinition {
	return &CSSDefinition{
		Selector: selector,
		DefLine:  line,
		Point:    point - selector.RawOffset,
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
