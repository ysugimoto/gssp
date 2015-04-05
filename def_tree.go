package gssp

type CSSDefinitionTree struct {
	definitions []*CSSDefinition
}

func NewDefinitionTree() *CSSDefinitionTree {
	return &CSSDefinitionTree{
		definitions: make([]*CSSDefinition, 0),
	}
}

func (l *CSSDefinitionTree) AddDefinitionToChild(def *CSSDefinition) {
	l.GetLastChild().AddChild(def)
}

func (l *CSSDefinitionTree) AddDefinition(def *CSSDefinition) {
	l.definitions = append(l.definitions, def)
}

func (l *CSSDefinitionTree) Remains() (remains bool) {
	if len(l.definitions) > 0 {
		remains = true
	}

	return
}

func (l *CSSDefinitionTree) HasParent() (has bool) {
	if len(l.definitions) > 1 {
		has = true
	}

	return
}

func (l *CSSDefinitionTree) GetLastChild() *CSSDefinition {
	return l.definitions[len(l.definitions)-1]
}

func (l *CSSDefinitionTree) Remove() {
	l.definitions = l.definitions[0 : len(l.definitions)-1]
}
