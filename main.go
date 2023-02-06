package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ============================================================================================== //
// 1. Http router, mirip kyk mux sbnrnya, buat routing supaya bisa akses bbrp page/directory
func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Test message ping for http router get method")
	})

	server := http.Server{
		Addr:    "localhost:2020",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// ============================================================================================== //
