package datasets

import (
	"encoding/json"
	"os"
	"sync"
)

type name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type image struct {
	URL      string `json:"url"`
	Category string `json:"category"`
	// TODO : more fields? size?
}

type Datasets struct {
	Names  []name
	Images []image
}

// TODO: Expose this publicly so that user can have their own custom datasets.
func read(filename string, schema interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Couldnt open file " + filename)
	}
	err = json.Unmarshal(data, schema)
	if err != nil {
		panic(err)
	}

}

// TODO: Add concurrency to this so it concurrently reads the thingys
func PopulateAll() *Datasets {
	dataset := Datasets{}
	var wg sync.WaitGroup
	wg.Add(2)
	go read("datasets/data/names.json", &dataset.Names, &wg)
	go read("datasets/data/images.json", &dataset.Images, &wg)
	wg.Wait()
	return &dataset
}
