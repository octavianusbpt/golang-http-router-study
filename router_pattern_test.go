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
// 4. Buat parameter dlm httprouter ada 2 jenis, yg pertama itu named parameter spt dibawah
// Kalo mau ambil variabel dari pattern 1 per 1, pakai titik dua (:)
func TestRouterPatternNamedParameter(t *testing.T) {

	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("id")
		itemId := p.ByName("itemId")
		fmt.Fprint(w, "Product "+productId+" Item "+itemId)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/products/1/items/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 2", string(body))

}

// ============================================================================================== //
// Yg kedua AllParameter spt dibawah
// Kalo mau ambil semua variabel / item dlm url stlh titik tertentu pakai bintang (*)
func TestRouterPatternAllParameter(t *testing.T) {

	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Image : "+p.ByName("image"))
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))

}

// ============================================================================================== //
