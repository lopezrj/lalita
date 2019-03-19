package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getTablesHandler(w http.ResponseWriter, r *http.Request) {
	tables, err := store.GetTables()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&tables)
}

func getPersonasHandler(w http.ResponseWriter, r *http.Request) {
	personas, err := store.GetPersonas()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&personas)
}
