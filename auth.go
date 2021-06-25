package main

import (
	"net/http"

	"github.com/michaeliyke/chitchat/data"
)

// "github.com/michaeliyke/chitchat/data"
// Get /login
// Show the login page

func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// Get /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	err = user.Create()
	if err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", http.StatusFound)
}

// POST /authenicate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:  "_cookie",
			Value: session.Uuid,
			// You set HttpOnly to only allow HTTP or HTTPS to access the
			// cookie (and not other non-HTTP APIs like JavaScript).
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie) // Add the cookie to the response header
		http.Redirect(writer, request, "/", http.StatusFound)
	} else {
		http.Redirect(writer, request, "/login", http.StatusFound)
	}
}

// GET /logout
// Logs the suer out
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", http.StatusFound)
}
