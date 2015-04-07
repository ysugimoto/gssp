package main

import (
	"fmt"
	"github.com/ysugimoto/go-cliargs"
	"github.com/ysugimoto/gssp"
	"os"
)

func usage() {
	text := `=================================
GSSP: Go Style-Sheet Postprocessor
=================================
Usage:
    gssp [source_file ...] [options]

Options:
    -h, --help   : Show this help
    -p, --pretty : Pretty print JSON result
`
	fmt.Println(text)
	os.Exit(0)
}

func main() {
	args := cliarg.NewArguments()
	args.Alias("h", "help", false)
	args.Alias("p", "pretty", false)
	args.Alias("s", "stats", false)

	args.Parse()

	if help, _ := args.GetOptionAsBool("help"); help {
		usage()
	}

	if args.GetCommandSize() == 0 {
		fmt.Println("[ERROR] Source CSS file must be suppied.")
		usage()
	}

	result := gssp.NewParseResult(
		make([]*gssp.CSSDefinition, 0),
	)

	files := args.GetCommands()
	for _, file := range files {
		parser := gssp.NewParser()
		if parsed, err := parser.ParseFile(file); err != nil {
			fmt.Println(err.Error)
			os.Exit(1)
		} else {
			result.Merge(parsed)
		}
	}

	var out string
	if s, _ := args.GetOptionAsBool("stats"); s {
		out = analyze(result)
	} else {
		pretty, _ := args.GetOptionAsBool("pretty")
		if pretty {
			out = result.ToPrettyJSONString()
		} else {
			out = result.ToJSONString()
		}
	}
	fmt.Print(out)
}

func analyze(result *gssp.CSSParseResult) string {
	return ""
}
