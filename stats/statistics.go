package stats

import (
	"github.com/ysugimoto/gssp"
	"regexp"
	"sort"
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
	//ruleAnalysis := s.analyzeRules()
	//selectorAnalysis := s.analyzeSelectors()

	// TODO: immplement
	//declAnalysis := Declarations.Analyze(s.declarations)

	//analysis := make(map[string]interface{})
}

func (s *Stats) analyzeRules() Rules {
	rules := Rules{
		cssDeclarations: RuleInfo{},
	}

	for _, def := range s.rules {
		r := Rule{
			selector: []string{},
			count:    len(def.Rules),
		}
		for _, selector := range strings.Split(def.Selector.String(), ",") {
			r.selector = append(r.selector, strings.Trim(selector, " "))
		}
		rules.cssDeclarations = append(rules.cssDeclarations, r)
	}

	sort.Sort(sort.Reverse(rules.cssDeclarations))
	for _, d := range rules.cssDeclarations {
		rules.totalCssDeclarations += d.count
	}

	return rules
}

func (s *Stats) analyzeSelectors() Selectors {
	ret := Selectors{
		identifiers: SelectorIdentifiers{},
	}

	// regexes
	attributeRegex := regexp.MustCompile("\\[.+\\]$")
	pseudoRegex := regexp.MustCompile("\\s?([>\\+|~])\\s?")
	whiteSpaceRegex := regexp.MustCompile("\\s+")
	splitCountRegex := regexp.MustCompile("\\s|>|\\+|~|:|[\\w\\]]\\.|[\\w\\]]#|\\[")

	for _, selector := range s.selectors {
		if strings.Contains(selector, "#") {
			ret.idSelectors++
		}

		if strings.Contains(selector, "*") {
			ret.universalSelectors++
		}

		sel := strings.Trim(selector, " ")

		if attributeRegex.MatchString(sel) {
			ret.unqualifiedAttributeSelectors++
		}

		// TODO: javascript/user-specified hook

		trimmed := pseudoRegex.ReplaceAllString(selector, "$1")
		trimmed = whiteSpaceRegex.ReplaceAllString(selector, " ")
		splitted := splitCountRegex.Split(trimmed, -1)
		ret.identifiers = append(ret.identifiers, SelectorIdentifier{
			selector: selector,
			count:    len(splitted),
		})
	}

	for _, ident := range ret.identifiers {
		ret.totalIdentifiers += ident.count
	}

	sort.Sort(ret.identifiers)

	return ret

}
