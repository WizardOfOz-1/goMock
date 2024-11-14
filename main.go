package main

import (
	"github.com/WizardOfOz-1/goMock/datasets"
	"github.com/WizardOfOz-1/goMock/parse"
)

func main() {
	dataset := datasets.PopulateAll()
	schema := parse.ParseInput()
	dataset.Generate(50, schema)
}
