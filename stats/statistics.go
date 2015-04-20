package stats

import (
	"github.com/ysugimoto/gssp"
	"strings"
)

type Stats struct {
	rules        []*gssp.CSSDefinition
	selectors    []string
	declarations []*gssp.CSSRule
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

func NewStats(parsedData *gssp.CSSParseResult) *Stats {

	s := &Stats{
		rules:        []*gssp.CSSDefinition{},
		selectors:    []string{},
		declarations: []*gssp.CSSRule{},
		cssString:    parsedData.RawBytes,
	}

	for _, def := range parsedData.Getdata() {
		s.divide(def)
	}

	return s
}

func (s *Stats) divide(def *gssp.CSSDefinition) {
	// rule
	s.rules = append(s.rules, def)

	// selectors
	for _, selector := range strings.Split(def.Selector.String(), ",") {
		s.selectors = append(s.selectors, strings.Trim(selector, " "))
	}

	// declarations
	for _, rule := range def.Rules {
		s.declarations = append(s.declarations, rule)
	}

	// has nested defintions?
	for _, control := range def.Controls {
		// call recursive
		s.divide(control)
	}
}

func (s *Stats) Analyze() {
	//ruleAnalysis := Rules.Analyze(s.rules)
	//selectorAnalysis := Selectors.Analyze(s.selectors)
	//declAnalysis := Declarations.Analyze(s.declarations)

	//analysis := make(map[string]interface{})
}
