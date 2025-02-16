package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/pat"
)

type Server struct {
	RoomList []*Room `json:"room_list"`
	Mutex    sync.RWMutex
}

type Room struct {
	ID string `json:"id"`
	// Clients  []*Client  `json:"clients"`
	// Messages []*Message `json:"messages"`
}

func main() {
	mux := pat.New()

	server := NewServer()

	mux.Get("/rooms", server.GetRooms)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":8080", n)
}

func NewServer() *Server {
	return &Server{
		RoomList: make([]*Room, 0),
	}
}

func (s *Server) GetRooms(w http.ResponseWriter, r *http.Request) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	json.NewEncoder(w).Encode(s.RoomList)
}
