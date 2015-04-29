package stats

import (
	"regexp"
	"strconv"
)

type Declarations struct {
	dataUriSize        int
	importantKeywords  int
	floatProperties    int
	uniqueFontSizes    FontSizes
	uniqueFontFamilies FontFamilies
	uniqueColors       Colors
	properties         Properties
}

var numRegex = regexp.MustCompile("[^0-9\\.]")

type FontSizes []string

func (fs FontSizes) Len() int {
	return len(fs)
}

func (fs FontSizes) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func (fs FontSizes) Less(i, j int) bool {
	var left, right float64

	left, _ = strconv.ParseFloat(numRegex.ReplaceAllString(fs[i], ""), 32)
	right, _ = strconv.ParseFloat(numRegex.ReplaceAllString(fs[j], ""), 32)

	return left < right
}

type FontFamilies []string

func (ff FontFamilies) Len() int {
	return len(ff)
}

func (ff FontFamilies) Swap(i, j int) {
	ff[i], ff[j] = ff[j], ff[i]
}

func (ff FontFamilies) Less(i, j int) bool {
	return ff[i] < ff[j]
}

type Colors []string

func (c Colors) Len() int {
	return len(c)
}

func (c Colors) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Colors) Less(i, j int) bool {
	return c[i] < c[j]
}

type Property struct {
	property string
	count    int
}

type Properties []Property

func (p Properties) Len() int {
	return len(p)
}

func (p Properties) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Properties) Less(i, j int) bool {
	return p[i].count < p[j].count
}
