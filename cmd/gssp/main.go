package main

import (
	"fmt"
	"github.com/ysugimoto/go-cliargs"
	"github.com/ysugimoto/gssp"
	"io/ioutil"
	"os"
	"path/filepath"
)

func usage() {
	text := `=================================
GSSP: Go Style-Sheet Postprocessor
=================================
Usage:
    gssp [source_file] [options]

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

	args.Parse()

	if help, _ := args.GetOptionAsBool("help"); help {
		usage()
	}

	file, exists := args.GetCommandAt(1)
	if !exists {
		fmt.Println("[ERROR] Source CSS file must be suppied.")
		usage()
	}
	absPath, _ := filepath.Abs(file)
	if _, err := os.Stat(absPath); err != nil {
		fmt.Println("[ERROR] Source CSS file " + file + " is not exists.")
		os.Exit(1)
	}

	prettyPrint, _ := args.GetOptionAsBool("pretty")
	result := execute(absPath, prettyPrint)

	fmt.Print(result)
}

func execute(path string, pretty bool) string {
	parser := gssp.NewParser()

	fp, _ := os.Open(path)
	buffer, _ := ioutil.ReadAll(fp)

	defer func() {
		fp.Close()
	}()

	result := parser.Parse(buffer)
	if pretty {
		return result.ToPrettyJSONString()
	} else {
		return result.ToJSONString()
	}
}
