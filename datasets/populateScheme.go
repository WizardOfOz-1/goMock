package datasets

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func (d *Datasets) getRandom(thing string) string {
	switch thing {
	case "name":
		return fmt.Sprintf("%s %s", d.Names[rand.Intn(len(d.Names))].First, d.Names[rand.Intn(len(d.Names))].Last)
	case "int":
		return strconv.Itoa(rand.Intn(69))
	case "img":
		return "https://www.yomama.com"
	case "bool":
		if rand.Intn(2) == 1 {
			return "true"
		}
		return "false"
	}
	return ""
}

func (d *Datasets) Generate(Times int, schema string) {
	for range Times {
		r := strings.NewReplacer("~name~", d.getRandom("name"), "~int~", d.getRandom("int"), "~image~", d.getRandom("img"), "~bool~", d.getRandom("bool"))
		fmt.Println(r.Replace(schema))
	}
}
