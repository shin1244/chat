package main

import (
	"net/http"
	"web/app"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/pat"
)

func main() {
	mux := pat.New()
	server := app.NewServer()

	mux.Post("/rooms", server.CreateRoom)
	mux.Get("/rooms", server.GetRooms)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":8080", n)
}
