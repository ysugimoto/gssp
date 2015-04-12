package css_test

import (
	"bytes"
	. "github.com/r7kamura/gospel"
	"github.com/ysugimoto/gssp"
	"io/ioutil"
	"os"
	"testing"
)

var cssList = []string{
	"./cases/atrule-decls",
	"./cases/atrule-no-semicolon",
	"./cases/colun-seector",
	"./cases/escape",
	"./cases/important",
	"./cases/rule-no-semicolon",
	"./cases/atrule-no-space",
	"./cases/comments",
	"./cases/extends",
	"./cases/prop-hacks",
	"./cases/selector",
	"./cases/atrule-empty",
	"./cases/atrule-params",
	"./cases/decls",
	"./cases/function",
	"./cases/quotes",
	"./cases/semicolons",
	"./cases/atrule-no-params",
	"./cases/atrule-rules",
	"./cases/empty",
	"./cases/ie-progid",
	"./cases/raw-decl",
	"./cases/tab",
}

func TestParser(t *testing.T) {

	Describe(t, "Test Cases From PostCSS", func() {

		for _, file := range cssList {
			Context(file+".css parse test", func() {
				It(file+".css should be same "+file+".json", func() {
					parser := gssp.NewParser()
					cssfp, _ := os.Open(file + ".css")
					jsonfp, _ := os.Open(file + ".json")

					defer func() {
						cssfp.Close()
						jsonfp.Close()
					}()

					data, _ := ioutil.ReadAll(cssfp)
					jsonData, _ := ioutil.ReadAll(jsonfp)
					actual := parser.Parse(data).ToPrettyJSONString()

					// Trim trail \n for vim, ...
					expect := string(bytes.TrimRight(jsonData, "\n"))

					Expect(actual).To(Equal, expect)
				})
			})
		}
	})
}
