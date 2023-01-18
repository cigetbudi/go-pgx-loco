package main

import (
	"bytes"
	"encoding/json"
	"go-pgx-loco/api"
	"go-pgx-loco/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRegister(t *testing.T) {
	r := api.InitRoutes()
	r.POST("/auth/registerTest", api.Register)
	u := model.User{
		Username: "usernameTest2",
		Password: "passwordTest2",
		Fullname: "Full NameTest2",
		Email:    "test2@mail.com",
		DOB:      "1990-01-02",
	}
	jsonValue, _ := json.Marshal(u)
	req, _ := http.NewRequest("POST", "/auth/registerTest", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
