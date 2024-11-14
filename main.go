package main

import (
	"fmt"

	"github.com/WizardOfOz-1/goMock/datasets"
	"github.com/WizardOfOz-1/goMock/parse"
)

/*
TODO:
    - More Data Sets
    - More Flexible Fields ( supporting lengths? )
    - Allow Users To Easily Add Their Own Custom Datasets?
    - Have a http handler that actually returns the response
*/
func main() {
	dataset := datasets.PopulateAll()
    fmt.Print(dataset.Images)
	schema := parse.ParseInput()
	dataset.Generate(50, schema)
}
