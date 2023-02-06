package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var folderResources embed.FS

// ============================================================================================== //
// 5. Serv file => Untuk get suatu file, tp pake httprouter
func TestServeFile(t *testing.T) {

	router := httprouter.New()
	// Buat masuk ke sub directory, krn kita gk mau ada keliatan folder (resources/) di dlm url
	// kita
	directory, _ := fs.Sub(folderResources, "resources")
	// NOTES: HARUS PAKE /*filepath => SUDAH DI HARDCODE BAWAAN HTTPROUTER
	router.ServeFiles("/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "This is a test message for serve file module in golang http router", string(body))

}

func TestServeFileSecondFile(t *testing.T) {

	router := httprouter.New()
	// Buat masuk ke sub directory, krn kita gk mau ada keliatan folder (resources/) di dlm url
	// kita
	directory, _ := fs.Sub(folderResources, "resources")
	// NOTES: HARUS PAKE /*filepath => SUDAH DI HARDCODE BAWAAN HTTPROUTER
	router.ServeFiles("/octaDummy/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2020/octaDummy/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Dummy text", string(body))

}

// ============================================================================================== //
