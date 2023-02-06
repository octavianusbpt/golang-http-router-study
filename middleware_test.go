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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive request")
	middleware.Handler.ServeHTTP(w, r)
}

// ============================================================================================== //
// 9. HttpRoter khusus buat router saja. Kalau ingin pakai middleware buat manual, ini contoh
// buatnya
func TestMiddleware(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware Test")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware Test", string(body))

}

// ============================================================================================== //
