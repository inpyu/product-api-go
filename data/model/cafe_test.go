package model

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCafeDeserializeFromJSON(t *testing.T) {
	c := Cafes{}

	err := c.FromJSON(bytes.NewReader([]byte(cafesData)))
	assert.NoError(t, err)

	assert.Len(t, c, 2)
	assert.Equal(t, 1, c[0].ID)
	assert.Equal(t, 2, c[1].ID)
}

func TestCafesSerializesToJSON(t *testing.T) {
	c := Cafes{
		Cafe{ID: 1, Name: "test", Address: "test"},
	}

	d, err := c.ToJSON()
	assert.NoError(t, err)

	cd := make([]map[string]interface{}, 0)
	err = json.Unmarshal(d, &cd)
	assert.NoError(t, err)

	assert.Equal(t, float64(1), cd[0]["id"])
	assert.Equal(t, "test", cd[0]["name"])
	assert.Equal(t, float64(120.12), cd[0]["address"])
}

var cafesData = `
[
	{
		"id": 1,
		"name": "StarBucks",
		"address": "seoul"
	},
	{
		"id": 2,
		"name": "Edia",
		"address": "seoul
	}
]
`