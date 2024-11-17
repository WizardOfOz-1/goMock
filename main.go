package main

import (
	"fmt"

	"github.com/WizardOfOz-1/Mockery/server"
	"github.com/WizardOfOz-1/Mockery/sets"
	"github.com/WizardOfOz-1/Mockery/utils"
)

func main() {
	dataset := sets.NewEmptySet()
	schema := utils.ReadFile("tests/userSet_test.json")
	dataset.ParseDataSet("tests/userSet.json")
	for range 10 {
		fmt.Println(dataset.ParseInput(schema))
	}
	server.StartServer(":3000", schema, *dataset, 5)
	// schema := utils.ReadFile("tests/schema.json")
	// Strings.replace everything that starts with ~...~ with the generated prefixes
	// fmt.Println(string(schema))

}
