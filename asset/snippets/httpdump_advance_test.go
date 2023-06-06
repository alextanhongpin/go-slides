package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alextanhongpin/core/test/testutil"
	"github.com/stretchr/testify/assert"
)

type loginUsecase interface {
	Login(ctx context.Context, email, password string) (token string, err error)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type LoginController struct {
	usecase loginUsecase
}

func (ctrl *LoginController) PostLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	token, err := ctrl.usecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res := LoginResponse{
		AccessToken: token,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type mockLoginUsecase struct {
	email    string
	password string
	token    string
	err      error
	dirty    bool
}

func (uc *mockLoginUsecase) Login(ctx context.Context, email, password string) (string, error) {
	if uc.dirty {
		panic("called more than once")
	}

	uc.email = email
	uc.password = password
	uc.dirty = true

	return uc.token, uc.err
}

func TestLoginHandlerSuccess(t *testing.T) {
	uc := &mockLoginUsecase{token: "token-xyz"}
	ctrl := &LoginController{usecase: uc}

	b := []byte(`{
			"email": "john.appleseed@mail.com",
			"password": "s3cur3p@$$w0rd"
		}`)
	r := httptest.NewRequest("POST", "/login", bytes.NewReader(b))
	testutil.DumpHTTP(t, r, ctrl.PostLogin)

	assert := assert.New(t)
	assert.Equal("john.appleseed@mail.com", uc.email)
	assert.Equal("s3cur3p@$$w0rd", uc.password)
	assert.ErrorIs(uc.err, uc.err)
}

func TestLoginHandlerFailed(t *testing.T) {
	err := errors.New("user not found")
	uc := &mockLoginUsecase{err: err}
	ctrl := &LoginController{usecase: uc}

	b := []byte(`{
			"email": "john.appleseed@mail.com",
			"password": "s3cur3p@$$w0rd"
		}`)
	r := httptest.NewRequest("POST", "/login", bytes.NewReader(b))
	testutil.DumpHTTP(t, r, ctrl.PostLogin)

	assert := assert.New(t)
	assert.Equal("john.appleseed@mail.com", uc.email)
	assert.Equal("s3cur3p@$$w0rd", uc.password)
	assert.ErrorIs(uc.err, uc.err)
}
