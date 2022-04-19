package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Store struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r*http.Request) {
	var store Store

	json.NewDecoder(r.Body).Decode(&store)
	fmt.Fprintf(w, "%s", store.Name)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8002", nil)
}
