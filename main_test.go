package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	res := map[string]string{}
	app := setupApp()
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	assert.Equal(t, 200, resp.StatusCode, "HTTP Code should be equal to 200")
	assert.Equal(t, "GO-MINOES", res["message"], "message should be equal to GO-MINOES")
}
