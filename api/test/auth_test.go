package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/mhdiiilham/gominoes/api/app"
	"github.com/mhdiiilham/gominoes/entity/user"
	mock "github.com/mhdiiilham/gominoes/mocks"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func TestRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urMock := mock.NewMockRepository(ctrl)
	pwdMockd := mock.NewMockHasher(ctrl)
	userManager := user.NewManager(urMock, pwdMockd)
	jwtMock := mock.NewMockTokenService(ctrl)
	managers := &app.Managers{
		UserManager: userManager,
	}

	res := map[string]string{}
	v, trans := app.SetupValidator()
	r := app.SetupApp(managers, jwtMock, v, trans)
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := r.Test(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	assert.Equal(t, 200, resp.StatusCode, "HTTP Code should be equal to 200")
	assert.Equal(t, "GO-MINOES", res["message"], "message should be equal to GO-MINOES")
}

func TestRegisterAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urMock := mock.NewMockManager(ctrl)
	pwdMock := mock.NewMockHasher(ctrl)
	userManager := user.NewManager(urMock, pwdMock)
	managers := &app.Managers{
		UserManager: userManager,
	}

	v, trans := app.SetupValidator()
	jwtMock := mock.NewMockTokenService(ctrl)
	r := app.SetupApp(managers, jwtMock, v, trans)

	testCases := []struct {
		Name    string
		Code    int
		Msg     string
		Payload user.User
	}{
		{
			Name:    "Register success",
			Code:    http.StatusCreated,
			Msg:     "Success created new user",
			Payload: generateUser(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			pwdMock.EXPECT().Hash(gomock.Any()).Return(tc.Payload.Password).Times(1)
			urMock.EXPECT().Register(tc.Payload).Return(&tc.Payload, nil).Times(1)
			jwtMock.EXPECT().Generate(&tc.Payload).Times(1)

			res := map[string]string{}
			payloadJSON, _ := json.Marshal(tc.Payload)
			bodyReader := strings.NewReader(string(payloadJSON))
			req := httptest.NewRequest("POST", "/api/auth/registrations", bodyReader)
			req.Header.Set("Content-Type", "application/json")
			resp, _ := r.Test(req)
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &res)
			assert.Equal(t, tc.Code, resp.StatusCode, "HTTP Status Code should be equal to created")
			assert.Equal(t, tc.Msg, res["message"], fmt.Sprintf("msg should be equal to %v", tc.Msg))
		})
	}
}

func generateUser() user.User {
	return user.User{
		Fullname: faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}
