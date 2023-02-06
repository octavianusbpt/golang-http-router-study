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
// 8. Kl user menggunakan method yg beda dr yg kita tentukan. Contoh: kita ingin menggunakan method
// GET pada suatu halaman, tetapi user malah mengirimkan method POST, maka httprouter sbnrnya
// sudah secara otomatis mengarahkan ke halaman error not allowed. Ini kita lakukan kl mau custom page
// error-nya
func TestMethodNotAllowed(t *testing.T) {

	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Method yang digunakan tidak sesuai >:(")
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Method POST")
	})

	// Manggil method get, HARUSNYA panggil method POST
	request := httptest.NewRequest(http.MethodPut, "http://localhost:2020/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Method yang digunakan tidak sesuai >:(", string(body))

}

// ============================================================================================== //
