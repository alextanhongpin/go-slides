package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alextanhongpin/core/test/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute business logic here ...
	u := User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		CreatedAt: time.Now(),
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestRegisterHandler(t *testing.T) {
	b := []byte(`{
		"email": "john.appleseed@mail.com",
		"password": "s3cur3p@$$w0rd",
		"name": "John Appleseed"
	}`)
	r := httptest.NewRequest("POST", "/register", bytes.NewReader(b))
	testutil.DumpHTTP(t, r, registerHandler,
		testutil.IgnoreFields("id", "createdAt"),
		testutil.InspectBody(func(payload []byte) {
			var u User
			err := json.Unmarshal(payload, &u)
			assert := assert.New(t)
			assert.Nil(err)

			id, err := uuid.Parse(u.ID)
			assert.Nil(err)
			assert.True(id != uuid.Nil)
		}),
		testutil.InspectHeaders(func(h http.Header) {
			contentType := h.Get("Content-Type")
			assert := assert.New(t)
			assert.Equal("text/plain; charset=utf-8", contentType)
		}),
	)
}
