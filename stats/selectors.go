package stats

type Selectors struct {
	idSelectors                   int
	universalSelectors            int
	unqualifiedAttributeSelectors int
	javascriptSpecificSelectors   int
	userSpecifiedSelectors        int
	totalIdentifiers              int
	identifiers                   SelectorIdentifiers
}

type SelectorIdentifier struct {
	selector string
	count    int
}

type SelectorIdentifiers []SelectorIdentifier

func (s SelectorIdentifiers) Len() int {
	return len(s)
}

func (s SelectorIdentifiers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SelectorIdentifiers) Less(i, j int) bool {
	return s[i].count < s[j].count
}
