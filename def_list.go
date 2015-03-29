package css

type CSSDefinitionList struct {
	definitions []*CSSDefinition
}

func NewDefinitionList() *CSSDefinitionList {
	return &CSSDefinitionList{
		definitions: make([]*CSSDefinition, 0),
	}
}

func (l *CSSDefinitionList) AddDefinitionToChild(def *CSSDefinition) {
	l.GetLastChild().AddChild(def)
}

func (l *CSSDefinitionList) AddDefinition(def *CSSDefinition) {
	l.definitions = append(l.definitions, def)
}

func (l *CSSDefinitionList) Remains() (remains bool) {
	if len(l.definitions) > 0 {
		remains = true
	}

	return
}

func (l *CSSDefinitionList) HasParent() (has bool) {
	if len(l.definitions) > 1 {
		has = true
	}

	return
}

func (l *CSSDefinitionList) GetLastChild() *CSSDefinition {
	return l.definitions[len(l.definitions)-1]
}

func (l *CSSDefinitionList) Remove() {
	l.definitions = l.definitions[0 : len(l.definitions)-1]
}
