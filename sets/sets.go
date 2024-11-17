package sets

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/WizardOfOz-1/Mockery/utils"
)

type DataSet struct {
	File     string
	Triggers map[string][]string
}

func NewEmptySet() *DataSet {
	return new(DataSet)
}

func (d *DataSet) ParseDataSet(path string) {
	var schema map[string]interface{}
	err := json.Unmarshal(utils.ReadFile(path), &schema)
	if err != nil {
		panic("Failed to unmarshal JSON: " + err.Error())
	}
	// Convert map[string]interface{} to map[string][]string
	mapped := make(map[string][]string)
	for key, val := range schema {
		array, ok := val.([]interface{})
		if !ok {
			panic("Please ensure that the dataset is in the form of a trigger followed by an array of strings")
		}
		for _, item := range array {
			str, ok := item.(string)
			if !ok {
				panic("Found a non-string item in the array")
			}
			triggerParsed := fmt.Sprintf("~%s~", key)
			mapped[triggerParsed] = append(mapped[triggerParsed], str)
		}
	}
	d.File = path
	d.Triggers = mapped
}

// Get Returns a randomly selected value for that trigger
func (d *DataSet) GetValue(trigger string) string {
	v, ok := d.Triggers[trigger]
	if !ok {
		return "No Data For That Trigger"
	}
	return v[rand.Intn(len(v))]
}

func (d *DataSet) ParseInput(schema []byte) string {
	re := regexp.MustCompile(`~[^~]+~`)
	matches := re.FindAll(schema, -1)
	var stringMatches []string
	for _, val := range matches {
		stringMatches = append(stringMatches, string(val), d.GetValue(string(val)))
	}
	replacer := strings.NewReplacer(stringMatches...)
	return replacer.Replace(string(schema))

}
