package main

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	app := newTestApplication(t)

	mux := app.mount()
	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/150", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := executeReq(req, mux)
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should allow aunthenticated requests", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/v1/users/150", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := executeReq(req, mux)
		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}
