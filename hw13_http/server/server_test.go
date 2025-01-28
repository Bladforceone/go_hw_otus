package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bladforceone/go_hw_otus/hw13_http/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	t.Run("Рабочий запрос", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/hello", nil)
		rr := httptest.NewRecorder()

		hello(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "Hello world!", rr.Body.String())
	})
	t.Run("Неправильный метод запроса", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/v1/hello", nil)
		rr := httptest.NewRecorder()

		hello(rr, req)

		var resp map[string]string
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		require.NoError(t, err)
		assert.Equal(t, "метод не разрешен", resp["error"])
	})
}

func TestGetUserHandler(t *testing.T) {
	t.Run("Рабочий GET-запрос", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/getUser", nil)
		rr := httptest.NewRecorder()
		getUser(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var user types.User
		err := json.Unmarshal(rr.Body.Bytes(), &user)
		require.NoError(t, err)
		assert.Equal(t, types.CreateExampleUser(), user)
	})

	t.Run("Неправильный метод", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/v1/getUser", nil)
		rr := httptest.NewRecorder()
		getUser(rr, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	})
}
