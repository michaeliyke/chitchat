package main

import "net/http"

func main() {
	mux := http.NewServeMux()                     // The multiplexer
	files := http.FileServer(http.Dir("/public")) // Handler to serve files
	// For all request URLs starting with /static/, strip off /static/ from the URL
	mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}
	server.ListenAndServe()
}
