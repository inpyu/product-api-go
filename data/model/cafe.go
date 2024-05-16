package model

import (
	"database/sql"
	"encoding/json"
	"io"
)

type Cafes []Cafe

// FromJSON serializes data from json
func (c *Cafes) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}

// ToJSON converts the collection to json
func (c *Cafes) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// Coffee defines a coffee in the database
type Cafe struct {
	ID          int            `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Address     string         `db:"address" json:"address"`
	Description string         `db:"description" json:"description"`
	Image       string         `db:"image" json:"image"`
	CreatedAt   string         `db:"created_at" json:"-"`
	UpdatedAt   string         `db:"updated_at" json:"-"`
	DeletedAt   sql.NullString `db:"deleted_at" json:"-"`
}

func (c *Cafe) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}

// ToJSON converts the collection to json
func (c *Cafe) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}
