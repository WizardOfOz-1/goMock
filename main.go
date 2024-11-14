package main

import (
	"fmt"
	"strings"

	"github.com/WizardOfOz-1/goMock/datasets"
	"github.com/WizardOfOz-1/goMock/parse"
)

func main() {
	dataset := datasets.PopulateAll()
	schema := parse.ParseInput()
	print(strings.ReplaceAll(schema, "~name~", fmt.Sprintf("%s %s", dataset.Names[0].First, dataset.Names[0].Last)))
}
