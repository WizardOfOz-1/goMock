package datasets

import (
	"encoding/json"
	"os"
)

type name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type image struct {
	URL string `json:"url"`
	// TODO : more fields? size?
}

type Datasets struct {
	Names  []name
	Images []image
}

func read(filename string, schema interface{}) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Couldnt open file " + filename)
	}
	err = json.Unmarshal(data, schema)
	if err != nil {
		panic(err)
	}

}

func PopulateAll() *Datasets {
	dataset := Datasets{}
	read("datasets/names.json", &dataset.Names)
	return &dataset
}
