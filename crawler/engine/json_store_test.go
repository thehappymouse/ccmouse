package engine

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGirlStore_WriteDisk(t *testing.T) {
	o := map[string]interface{} {
		"id":"4",
		"name": "go_test",
	}
	store := JsonStore{JsonPath: "test.json"}
	store.Set("id1", o)
	store.WriteDisk()

	store2 := CreateJsonStore("test.json")

	assert.Equal(t, store2.Get("id1"), o)
}
