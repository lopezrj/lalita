package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Store will have two methods, to add a new bird,
// and to get all existing birds
// Each method returns an error, in case something goes wrong
type Store interface {
	GetTables() ([]*Table, error)
	GetPersonas() ([]*Persona, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type dbStore struct {
	db *sql.DB
}

// Table is table
type Table struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	TblName  string `json:"tbl_name"`
	Rootpage string `json:"rootpage"`
	SQL      string `json:"-"`
}

// Persona is a persona
type Persona struct {
	Apellido1       string `json:"apellido1"`
	Apellido2       string `json:"apellido2"`
	Nombre1         string `json:"nombre1"`
	Nombre2         string `json:"nombre2"`
	Sexo            string `json:"sexo"`
	FechaNacimiento string `json:"fecha_nacimiento"`
}

func (store *dbStore) GetTables() ([]*Table, error) {
	sqlStatement := `SELECT * FROM sqlite_master WHERE type='table';`
	rows, err := store.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := []*Table{}
	for rows.Next() {
		table := &Table{}
		if err := rows.Scan(
			&table.Type,
			&table.Name,
			&table.TblName,
			&table.Rootpage,
			&table.SQL); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (store *dbStore) GetPersonas() ([]*Persona, error) {
	sqlStatement := `SELECT c_apellido1, c_apellido2, c_nombre1, c_nombre2, c_sexo, d_fecha_nacimiento FROM SIE_C_PERSONA limit 10`
	rows, err := store.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	personas := []*Persona{}
	for rows.Next() {
		persona := &Persona{}
		if err := rows.Scan(
			&persona.Apellido1,
			&persona.Apellido2,
			&persona.Nombre1,
			&persona.Nombre2,
			&persona.Sexo,
			&persona.FechaNacimiento); err != nil {
			return nil, err
		}
		personas = append(personas, persona)
	}
	return personas, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
InitStore We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	store = s
}
