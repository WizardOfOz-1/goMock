package parse

import (
	"os"
)

// Possible port this to take input from the user via a web-browser?
func ParseInput() string {
	schema, err := os.ReadFile("tests/schema.json")
	if err != nil {
		panic("Couldn't read schema")
	}
	return string(schema)

}
