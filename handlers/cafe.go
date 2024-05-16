package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp-demoapp/product-api-go/data"
	"github.com/hashicorp-demoapp/product-api-go/data/model"
	"github.com/hashicorp/go-hclog"
)

// Coffee -
type Cafe struct {
	con data.Connection
	log hclog.Logger
}

// NewCafe
func NewCafe(con data.Connection, l hclog.Logger) *Cafe {
	return &Cafe{con, l}
}

func (c *Cafe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Cafe")

	vars := mux.Vars(r)

	var cafeID *int

	if vars["id"] != "" {
		cId, err := strconv.Atoi(vars["id"])
		if err != nil {
			c.log.Error("Cafe provided could not be converted to an integer", "error", err)
			http.Error(rw, "Unable to list ingredients", http.StatusInternalServerError)
			return
		}
		cafeID = &cId
	}

	cofs, err := c.con.GetCafes(cafeID)
	if err != nil {
		c.log.Error("Unable to get products from database", "error", err)
		http.Error(rw, "Unable to list products", http.StatusInternalServerError)
		return
	}

	d, err := cofs.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert products to JSON", "error", err)
		http.Error(rw, "Unable to list products", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

// CreateCafe creates a new cafe
func (c *Cafe) CreateCafe(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Cafe | CreateCafe")

	body := model.Cafe{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	cafe, err := c.con.CreateCafe(body)
	if err != nil {
		c.log.Error("Unable to create new cafe", "error", err)
		http.Error(rw, fmt.Sprintf("Unable to create new cafe: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	d, err := cafe.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert cafe to JSON", "error", err)
		http.Error(rw, "Unable to create new cafe", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

func (c *Cafe) UpdateCafe(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Cafe | CreateCafe")

	vars := mux.Vars(r)

	body := model.Cafe{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	cafeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("cafeID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to delete order", http.StatusInternalServerError)
		return
	}

	cafe, err := c.con.UpdateCafe(cafeID, body)
	if err != nil {
		c.log.Error("Unable to update cafe", "error", err)
		http.Error(rw, fmt.Sprintf("Unable to update cafe: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	d, err := cafe.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert cafe to JSON", "error", err)
		http.Error(rw, "Unable to create new cafe", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

// DeleteOrder deletes a user order
func (c *Cafe) DeleteCafe(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Cafes | DeleteCafe")

	vars := mux.Vars(r)

	cafeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("cafeID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to delete order", http.StatusInternalServerError)
		return
	}

	err = c.con.DeleteCafe(cafeID)
	if err != nil {
		c.log.Error("Unable to delete cafe from database", "error", err)
		http.Error(rw, "Unable to delete cafe", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "%s", "Deleted cafe")
}
