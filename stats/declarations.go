package stats

type Declarations struct {
	dataUriSize        string
	importantKeywords  int
	floatProperties    int
	uniqueFontSizes    []string
	uniqueFontFamilies []string
	uniqueColors       []string
	properties         map[string]int
}
