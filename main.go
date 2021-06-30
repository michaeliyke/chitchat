package main

import (
	"net/http"
	"time"
)

func main() {

	p("Chitchat", version(), "started at ", config.Address)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static)) // Handler to serve files

	// For all request URLs starting with /static/, strip off /static/ from the URL
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Defined in index.go
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	// Defined in auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// Defined in thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:           config.Address,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
		Handler:        mux,
	}
	server.ListenAndServe()
}
