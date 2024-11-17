package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WizardOfOz-1/Mockery/sets"
)

func StartServer(port string, schema []byte, dataSet sets.DataSet, times int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var response []map[string]string
		for range times {
			jsonStr := dataSet.ParseInput(schema)
			var parsed map[string]string
			err := json.Unmarshal([]byte(jsonStr), &parsed)
			if err != nil {
				http.Error(w, fmt.Sprintf("Could not parse input: %+v", err), http.StatusInternalServerError)
				return
			}
			response = append(response, parsed)
		}
		jsonValue, err := json.Marshal(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not marshal response: %+v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonValue))
	})
	println("Started Server At: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(fmt.Sprintf("Couln't start server %+v", err))
	}
}
