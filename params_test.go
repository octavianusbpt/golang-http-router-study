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
// 3. Params merupakan parameter ketiga dr router.Handle, isinya adalah parameter di dlm url
// yg bersifat DINAMIS > inilah advantage httprouter dibandingkan dgn mux, http router ini bisa
// ambil info dr url secara dinamis, 1 lagi, di http router method yg mau kita pake harus ditentuin
// dr awal, kyk contoh dibawah pake router.GET, bisa juga pake yg lain kyk router.POST, router.PUT
// router.DELETE
func TestParams(t *testing.T) {

	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("id")
		fmt.Fprint(w, "Product "+productId)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))

}

// ============================================================================================== //
