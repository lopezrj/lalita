package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	connString := "file:/Users/lopezrj/Data/sie.sqlite3"
	db, err := sql.Open("sqlite3", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	InitStore(&dbStore{db: db})
	defer db.Close()

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.Handle("/tables", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(getTablesHandler)))
	r.HandleFunc("/personas", getPersonasHandler)

	fmt.Println("Successfully connected!")
	log.Printf("Server listening on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8080), r))
}
