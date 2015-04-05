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

## License

MIT

## Author

Yoshiaki Sugimoto


