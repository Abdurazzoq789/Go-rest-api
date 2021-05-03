package app

import (
	"github.com/gorilla/mux"
	"postapi/app/database"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/posts", a.GetPostHandler()).Methods("GET")
	a.Router.HandleFunc("/api/post/{id}", a.GetOnePostHandler()).Methods("GET")
	a.Router.HandleFunc("/api/post/{id}", a.UpdatePostHandler()).Methods("PUT")
	a.Router.HandleFunc("/api/post/delete/{id}", a.DeletePostHandler()).Methods("DELETE")
}
