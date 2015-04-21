package stats

type Rule struct {
	selector []string
	count    int
}

type Rules struct {
	totalCssDeclarations int
	cssDeclarations      RuleInfo
}

type RuleInfo []Rule

func (ri RuleInfo) Len() int {
	return len(ri)
}

func (ri RuleInfo) Swap(i, j int) {
	ri[i], ri[j] = ri[j], ri[i]
}

func (ri RuleInfo) Less(i, j int) bool {
	return ri[i].count < ri[j].count
}
