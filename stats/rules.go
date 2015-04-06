package stats

type Rule struct {
	selector []string
	count    int
}

type Rules struct {
	cssDeclarations []Rule
}
