package css

type CSSDefinitionList struct {
	definitions []*CSSDefinition
}

func NewDefinitionList() *CSSDefinitionList {
	return &CSSDefinitionList{
		definitions: make([]*CSSDefinition, 0),
	}
}

func (l *CSSDefinitionList) Add(def *CSSDefinition) {
	l.definitions = append(l.definitions, def)
}

func (l *CSSDefinitionList) Merge(defs []*CSSDefinition) {
	l.definitions = append(l.definitions, defs...)
}
func (l *CSSDefinitionList) Get() []*CSSDefinition {
	return l.definitions
}
