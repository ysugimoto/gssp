package css

import "regexp"
import "bytes"

var (
	leftRegex    = regexp.MustCompile("^([;\n\t ]*)")
	rightRegex   = regexp.MustCompile("[\n\t/\\* ]*$")
	spaceRegex   = regexp.MustCompile("[\n ]+")
	commentRegex = regexp.MustCompile("/\\*.*\\*/")
)

func parseBytes(data []byte) (before, value, after []byte, offset int) {
	size := len(data)
	data = commentRegex.ReplaceAll(data, []byte(""))
	left := leftRegex.FindSubmatchIndex(data)
	before = data[:left[1]]
	data = data[left[1]:]

	right := rightRegex.FindSubmatchIndex(data)
	value = data[:right[0]]
	after = data[right[0]:]

	offset = size - len(before)
	value = spaceRegex.ReplaceAll(value, []byte(" "))

	return
}

func isEmptyStack(stack []byte) (isEmpty bool) {
	if len(bytes.Trim(stack, "\r\n\t ")) == 0 {
		isEmpty = true
	}

	return
}
