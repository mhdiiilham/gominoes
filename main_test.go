package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestRegister(t *testing.T) {
	app := setupApp()

	testCases := []struct {
		name    string
		code    int
		msg     string
		payload map[string]string
	}{
		{
			name: "Register success",
			code: http.StatusCreated,
			msg:  "Success created new user",
			payload: map[string]string{
				"fullname": "Super Testing",
				"email":    "super@testing.com",
				"password": "superpassword",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := map[string]string{}
			payloadJSON, _ := json.Marshal(tc.payload)
			bodyReader := strings.NewReader(string(payloadJSON))
			req := httptest.NewRequest("POST", "/auth/registrations", bodyReader)
			resp, _ := app.Test(req)
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &res)
			assert.Equal(t, tc.code, res["code"], "HTTP Status Code should be equal to created")
			assert.Equal(t, tc.payload["message"], res["message"], fmt.Sprintf("msg should be equal to %s", tc.payload["message"]))
		})
	}
}
