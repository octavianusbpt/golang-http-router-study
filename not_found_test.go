package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// ============================================================================================== //
// 7. Kl user coba masuk ke url yg gk di handle sama router (url random), biasanya bakal diredirect
// ke halaman not found. nah disini kita bisa custom sndr halaman not foundnya mau di handle gmn
func TestNotFound(t *testing.T) {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Halaman tidak ditemukan")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/urlrandom", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Halaman tidak ditemukan", string(body))

}

// ============================================================================================== //
