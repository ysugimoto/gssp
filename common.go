package css

import "regexp"

var (
	leftRegex  = regexp.MustCompile("^([;\n\t ]*)")
	rightRegex = regexp.MustCompile("[\n\t ]*$")
)

func parseBytes(data []byte) (before, value, after []byte) {
	left := leftRegex.FindSubmatchIndex(data)
	before = data[:left[1]]
	data = data[left[1]:]

	right := rightRegex.FindSubmatchIndex(data)
	value = data[:right[0]]
	after = data[right[0]:]

	return
}
