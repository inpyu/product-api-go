package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
			c.log.Error("Cafe ID could not be converted to an integer", "error", err)
			http.Error(rw, "Invalid cafe ID", http.StatusInternalServerError)
			return
		}
		cafeID = &cId
	}

	cofs, err := c.con.GetCafes(cafeID)
	if err != nil {
		c.log.Error("Unable to get cafes from database", "error", err)
		http.Error(rw, "Unable to list cafes", http.StatusInternalServerError)
		return
	}

	var d []byte
	d, err = json.Marshal(cofs)

	if err != nil {
		c.log.Error("Unable to convert cafes to JSON", "error", err)
		http.Error(rw, "Unable to list cafes", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

// CreateCafe creates a new cafe
func (c *Cafe) CreateCafe(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Cafe | CreateCafe")

	var cafes []model.Cafe

	// 요청 본문을 읽고 출력합니다.
	reqBody, _ := io.ReadAll(r.Body)
	c.log.Info("Request Body", "body", string(reqBody))
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody)) // 요청 본문을 리셋합니다.

	err := json.NewDecoder(r.Body).Decode(&cafes)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	c.log.Info("Decoded Body", "body", cafes)

	// 단일 카페 객체로 처리
	cafe := cafes[0]

	createdCafe, err := c.con.CreateCafe(cafe)
	if err != nil {
		c.log.Error("Unable to create new cafe", "error", err)
		http.Error(rw, fmt.Sprintf("Unable to create new cafe: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	d, err := createdCafe.ToJSON()
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
