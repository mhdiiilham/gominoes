package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mhdiiilham/gominoes/api/app"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func TestRoot(t *testing.T) {
	res := map[string]string{}
	v, trans := app.SetupValidator()
	r := app.SetupApp(v, trans)
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := r.Test(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	assert.Equal(t, 200, resp.StatusCode, "HTTP Code should be equal to 200")
	assert.Equal(t, "GO-MINOES", res["message"], "message should be equal to GO-MINOES")
}

func TestRegister(t *testing.T) {
	v, trans := app.SetupValidator()
	r := app.SetupApp(v, trans)

	testCases := []struct {
		Name    string
		Code    int
		Msg     string
		Payload map[string]string
	}{
		{
			Name: "Register success",
			Code: http.StatusCreated,
			Msg:  "Success created new user",
			Payload: map[string]string{
				"fullname": "Super Testing",
				"email":    "super@testing.com",
				"password": "superpassword",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := map[string]string{}
			payloadJSON, _ := json.Marshal(tc.Payload)
			fmt.Println(string(payloadJSON))
			bodyReader := strings.NewReader(string(payloadJSON))
			req := httptest.NewRequest("POST", "/api/auth/registrations", bodyReader)
			req.Header.Set("Content-Type", "application/json")
			resp, _ := r.Test(req)
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &res)
			fmt.Println(res["code"])
			assert.Equal(t, tc.Code, resp.StatusCode, "HTTP Status Code should be equal to created")
			assert.Equal(t, tc.Msg, res["message"], fmt.Sprintf("msg should be equal to %v", tc.Msg))
		})
	}
}
