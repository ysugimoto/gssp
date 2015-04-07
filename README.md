# GSSP

Golang Style Sheet Postprocessor

## Installation

```
$ go get github.com/ysugimoto/gssp/cmd/gssp
```

## Usage

Run with `-h` option

```
$ gssp -h
=================================
GSSP: Go Style-Sheet Postprocessor
=================================
Usage:
    gssp [source_file] [options]

Options:
    -h, --help   : Show this help
    -p, --pretty : Pretty print JSON result
```

## Format

GSSP will parse your CSS file into selector-rule-value structures. Please see parse result format in testcases.

https://github.com/ysugimoto/gssp/blob/master/test/cases/decls.json

## Using with your Golang program

Import and create Parser and execute.

```Go
import "github.com/ysugimoto/gssp"

func main() {
    parser := gssp.NewParser()
    result := parser.Parse(css buffer as []byte)

    // Format JSON
    out := result.ToJSONString()
    // Format Pretty JSON
    out := result.ToPrettyJSONString()
    // Get parse Data
    data := result.Get()
    // .. and modify tree as you like
}
```

## License

MIT

## Author

Yoshiaki Sugimoto

## Thanks

[PostCSS](https://github.com/postcss/postcss): for test cases and way of thinking

