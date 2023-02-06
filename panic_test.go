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
// 6. Misal di web kita ada terjadi panic, ini cara handle-nya supaya tdk tb2 no response begitu
// saja
func TestPanicHandler(t *testing.T) {

	router := httprouter.New()

	// Disini bagian yg kita tambahin panic handler-nya
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic : ", i)
	}
	// Ini router.get biasa yg kita tambahin panic
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))

}

// ============================================================================================== //
