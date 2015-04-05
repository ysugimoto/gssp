package css

import "regexp"

var (
	leftRegex    = regexp.MustCompile("^([;\n\t ]*)")
	rightRegex   = regexp.MustCompile("[\n\t/\\* ]*$")
	spaceRegex   = regexp.MustCompile("[\n ]+")
	commentRegex = regexp.MustCompile("/\\*.+\\*/")
)

func parseBytes(data []byte) (before, value, after []byte) {
	data = commentRegex.ReplaceAll(data, []byte(""))
	left := leftRegex.FindSubmatchIndex(data)
	before = data[:left[1]]
	data = data[left[1]:]

	right := rightRegex.FindSubmatchIndex(data)
	value = data[:right[0]]
	after = data[right[0]:]

	value = spaceRegex.ReplaceAll(value, []byte(" "))

	return
}
