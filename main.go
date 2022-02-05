package main

import (
	"flag"
	// "fmt"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "The csv file contains all questions in format of question,answer")
	flag.Parse()
	_ = csvFilename
}
