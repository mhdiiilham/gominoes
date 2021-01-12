package api_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/gominoes/api/app"
	"github.com/mhdiiilham/gominoes/entity/user"
	mock "github.com/mhdiiilham/gominoes/mocks"
	"github.com/stretchr/testify/assert"
)

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

	payload := generateUser()

	pwdMock.EXPECT().Hash(gomock.Any()).Return(payload.Password).Times(1)
	urMock.EXPECT().Register(payload).Return(&payload, nil).Times(1)
	jwtMock.EXPECT().Generate(&payload).Times(1)

	res := map[string]string{}
	payloadJSON, _ := json.Marshal(payload)
	bodyReader := strings.NewReader(string(payloadJSON))
	req := httptest.NewRequest("POST", "/api/auth/registrations", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := r.Test(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "HTTP Status Code should be equal to CREATED")
	assert.Equal(t, "Success created new user", res["message"], "msg should be equal to Success created new user")
}

func TestRegisterAPIDuplicateEmail(t *testing.T) {
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

	payload := generateUser()

	pwdMock.EXPECT().Hash(gomock.Any()).Return(payload.Password).Times(1)
	urMock.EXPECT().Register(payload).Return(nil, errors.New("DUPLICATE EMAIL")).Times(1)

	res := map[string]string{}
	payloadJSON, _ := json.Marshal(payload)
	bodyReader := strings.NewReader(string(payloadJSON))
	req := httptest.NewRequest("POST", "/api/auth/registrations", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := r.Test(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	assert.Equal(t, http.StatusConflict, resp.StatusCode, "HTTP Status Code should be equal to CONFLICT")
	assert.Equal(t, "BAD REQUEST", res["message"], "msg should be equal to 'BAD REQUEST'")
}

func generateUser() user.User {
	return user.User{
		Fullname: faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}
