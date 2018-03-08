package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	var jsonBytes, err = json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Failed to marshal data: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonBytes))
}

func ExtractUuid(r *http.Request) string {
	return mux.Vars(r)["uuid"]
}
