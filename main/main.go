package main

import (
	"css"
	"fmt"
	"io/ioutil"
)

func main() {
	buffer, err := ioutil.ReadFile("./sample.css")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	parser := css.NewParser()

	fmt.Println(parser.Parse(buffer))
}
