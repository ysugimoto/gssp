package stats

import (
	"bytes"
	"compress/gzip"
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
	AverageOfIdentifier           float32
	AverageOfCohesion             float32
	MostIdentifier                int
	MostIdentifierSelector        string
	LowestCohesion                int
	LowersCohesionSelector        string
	TotalUniqueFontSizes          int
	UniqueFontSizes               FontSizes
	TotalUniqueColors             int
	UniqueColors                  Colors
	TotalUniqueFontFamilies       int
	UniqueFontFamilies            FontFamilies
	IdSelectors                   int
	UniversalSelectors            int
	UnqualifiedAttributeSelectors int
	JavaScriptSpecificSelectors   int
	ImportantKeywords             int
	FloatProperties               int
	PropertiesCount               Properties
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

func (s *Stats) Analyze() {
	ruleAnalysis := s.analyzeRules()
	selectorAnalysis := s.analyzeSelectors()
	declAnalysis := s.analyzeDeclarations()

	// TODO: support mutiple stylesheets
	s.StyleSheets = 1
	s.Size = len(s.cssString)
	s.DataUriSize = declAnalysis.dataUriSize
	s.RatioOfSataUriSize = declAnalysis.dataUriSize / s.Size
	s.GzippedSize = s.calculateGzippedSize(s.cssString)
	s.Rules = len(s.rules)
	s.Selectors = len(s.selectors)
	s.Simplicity = float32(s.Rules) / float32(s.Selectors)
	s.AverageOfIdentifier = float32(selectorAnalysis.totalIdentifiers) / float32(s.Selectors)

	if len(selectorAnalysis.identifiers) > 0 {
		s.MostIdentifierSelector = selectorAnalysis.identifiers[0].selector
	}

	s.AverageOfCohesion = float32(ruleAnalysis.totalCssDeclarations) / float32(s.Rules)

	if len(ruleAnalysis.cssDeclarations) > 0 {
		s.LowestCohesion = ruleAnalysis.cssDeclarations[0].count
		s.LowersCohesionSelector = strings.Join(ruleAnalysis.cssDeclarations[0].selector, ", ")
	}

	s.TotalUniqueFontSizes = len(declAnalysis.uniqueFontSizes)
	s.UniqueFontSizes = declAnalysis.uniqueFontSizes
	s.TotalUniqueColors = len(declAnalysis.uniqueColors)
	s.UniqueColors = declAnalysis.uniqueColors
	s.TotalUniqueFontFamilies = len(declAnalysis.uniqueFontFamilies)
	s.UniqueFontFamilies = declAnalysis.uniqueFontFamilies
	s.IdSelectors = selectorAnalysis.idSelectors
	s.UniversalSelectors = selectorAnalysis.universalSelectors
	s.UnqualifiedAttributeSelectors = selectorAnalysis.unqualifiedAttributeSelectors
	s.JavaScriptSpecificSelectors = selectorAnalysis.javascriptSpecificSelectors
	s.ImportantKeywords = declAnalysis.importantKeywords
	s.FloatProperties = declAnalysis.floatProperties
	s.PropertiesCount = declAnalysis.properties
}

func (s *Stats) calculateGzippedSize(css []byte) int {
	var w bytes.Buffer

	writer := gzip.NewWriter(&w)
	writer.Write(css)

	if err := writer.Flush(); err != nil {
		return 0
	}
	if err := writer.Close(); err != nil {
		return 0
	}
	return len(w.Bytes())
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

func (s *Stats) analyzeDeclarations() Declarations {
	ret := Declarations{
		uniqueFontSizes:    FontSizes{},
		uniqueFontFamilies: FontFamilies{},
		uniqueColors:       Colors{},
		properties:         Properties{},
	}

	importantRegex := regexp.MustCompile("(.+)!\\s+important")
	hexShortColorRegex := regexp.MustCompile("^#([0-9A-F])([0-9A-F])([0-9A-F])$")
	dataUriRegex := regexp.MustCompile("(data:image/[A-Za-z0-9,;\\+=/]+)")
	dataUris := []byte{}
	colors := []string{}
	props := map[string]int{}

	for _, decl := range s.declarations {
		prop := decl.Property
		value := decl.Value.Value
		if strings.Contains(value, "data:image") {
			di := []byte(dataUriRegex.ReplaceAllString(value, "$1"))
			dataUris = append(dataUris, di...)
		}
		if importantRegex.MatchString(value) {
			ret.importantKeywords += 1
		}
		if strings.Contains(prop, "float") {
			ret.floatProperties += 1
		}
		if strings.Contains(prop, "font-family") {
			ff := importantRegex.ReplaceAllString(value, "$1")
			ret.uniqueFontFamilies = append(ret.uniqueFontFamilies, strings.Trim(ff, " "))
		}
		if strings.Contains(prop, "font-size") {
			fs := importantRegex.ReplaceAllString(value, "$1")
			ret.uniqueFontSizes = append(ret.uniqueFontSizes, strings.Trim(fs, " "))
		}
		if strings.Trim(prop, " ") == "color" {
			c := importantRegex.ReplaceAllString(value, "$1")
			colors = append(colors, strings.Trim(strings.ToUpper(c), " "))
		}

		if _, ok := props[prop]; ok {
			props[prop] += 1
		} else {
			props[prop] = 1
		}
	}

	for _, c := range colors {
		if c == "TRANSPARENT" || c == "INHERIT" {
			continue
		}
		if hexShortColorRegex.MatchString(c) {
			c = hexShortColorRegex.ReplaceAllString(c, "#$1$1$2$2$3$3")
		}
		ret.uniqueColors = append(ret.uniqueColors, c)
	}

	for prop, count := range props {
		ret.properties = append(ret.properties, Property{
			property: prop,
			count:    count,
		})
	}

	ret.dataUriSize = len(dataUris)
	sort.Sort(ret.uniqueFontFamilies)
	sort.Sort(ret.uniqueFontSizes)
	sort.Sort(ret.properties)

	return ret
}
