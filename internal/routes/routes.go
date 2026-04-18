// Package routes contains the router
package routes

import (
	"net/http"
)

type Handling interface {
	Login(w http.ResponseWriter, r *http.Request)
	LoginPage(w http.ResponseWriter, r *http.Request)
	DashboardPage(w http.ResponseWriter, r *http.Request)
	HomePage(w http.ResponseWriter, r *http.Request)
	CreateAccount(w http.ResponseWriter, r *http.Request)
	NewAccount(w http.ResponseWriter, r *http.Request)
	NewPost(w http.ResponseWriter, r *http.Request)
	FeedPage(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type MiddleWare interface {
	Auth(next http.Handler) http.Handler
}

type Routes struct {
	Handling   Handling
	MiddleWare MiddleWare
}

func NewRouter(h Handling, m MiddleWare) *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.Handle("GET /", http.RedirectHandler("/dashboard", http.StatusMovedPermanently))
	mux.HandleFunc("GET /loginpage", h.LoginPage)
	mux.HandleFunc("POST /login", h.Login)
	mux.HandleFunc("GET /createaccount", h.CreateAccount)
	mux.HandleFunc("POST /newaccount", h.NewAccount)

	mux.Handle("GET /dashboard", m.Auth(http.HandlerFunc(h.DashboardPage)))
	mux.Handle("GET /home", m.Auth(http.HandlerFunc(h.HomePage)))
	mux.Handle("GET /feed", m.Auth(http.HandlerFunc(h.FeedPage)))

	mux.Handle("POST /post", m.Auth(http.HandlerFunc(h.NewPost)))
	mux.Handle("DELETE /deletepost", m.Auth(http.HandlerFunc(h.DeletePost)))

	return mux
}
