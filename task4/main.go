package main

import (
	"io/ioutil"
	"strconv"
)

func main() {
	inputData, err := ioutil.ReadFile("INPUT.TXT")
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(string(inputData))
	if err != nil {
		panic(err)
	}
	E := "2.7182818284590452353602875"

	var output string
	if n == 0 {
		output = "3"
	} else {
		output = E[:2+n]
	}

	err = ioutil.WriteFile("OUTPUT.TXT", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}
