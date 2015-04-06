package stats

import (
	"github.com/ysugimoto/gssp"
)

type Stats struct {
	rules        []*gssp.CSSRule
	selectors    []*gssp.CSSSelector
	declarations []*gssp.CSSDefinition
	cssString    []byte

	StyleSheets                   int
	Size                          int
	DataUriSize                   int
	RatioOfSataUriSize            int
	GzippedSize                   int
	Rules                         int
	Selectors                     int
	Simplicity                    float32
	MostIdentifier                int
	MostIdentifierSelector        string
	LowestCohesion                int
	LowersCohesionSelector        int
	TotalUniqueFontSizes          int
	UniqueFontSizes               string
	TotalUniqueColors             int
	UniqueColors                  string
	TotalUniqueFontFamilies       int
	UniqueFontFamilies            string
	IdSelectors                   int
	UniversalSelectors            int
	UnqualifiedAttributeSelectors int
	JavaScriptSpecificSelectors   int
	ImportantKeywords             int
	FloatProperties               int
	PropertiesCount               int
}

func NewStats(
	rules []*gssp.CSSRule,
	selectors []*gssp.CSSSelector,
	declarations []*gssp.CSSDefinition,
	cssString []byte) *Stats {

	return &Stats{
		rules:        rules,
		selectors:    selectors,
		declarations: declarations,
		cssString:    cssString,
	}
}

func (s *Stats) Analyze() {
	//ruleAnalysis := s.analyzeRules()
	//selectorAnalysis := s.analyzeSelectors()
	//declAnalysis := s.analyzeDeclarations()

	//analysis := make(map[string]interface{})
}

func (s *Stats) analyzeRules() Rules {
	return Rules{}
}

func (s *Stats) analyzeSelectors() Selectors {
	return Selectors{}
}

func (s *Stats) analyzeDeclarations() Declarations {
	return Declarations{}
}
