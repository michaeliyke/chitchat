package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_Login(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/login", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code Is %v", writer.Code)
	}

	body := writer.Body.String()
	if strings.Contains(body, "Sign in") == false {
		t.Fatal("Body does not contain Sign in")
	}
}

func Test_Exec(t *testing.T) {
	mux := http.NewServeMux()                     // The multiplexer
	files := http.FileServer(http.Dir("/public")) // Handler to serve files
	// For all request URLs starting with /static/, strip off /static/ from the URL
	log.Println(files)
	mux.Handle("/static", http.StripPrefix("/public", files))
}
