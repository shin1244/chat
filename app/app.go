package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type Room struct {
	ID string `json:"id"`
}

type Server struct {
	RoomList []*Room `json:"room_list"`
	Mutex    sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		RoomList: make([]*Room, 0),
	}
}

func (s *Server) CreateRoom(w http.ResponseWriter, r *http.Request) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	roomID := uuid.New().String()
	room := &Room{ID: roomID}

	s.RoomList = append(s.RoomList, room)
	fmt.Println(s.RoomList)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(room)
}

func (s *Server) GetRooms(w http.ResponseWriter, r *http.Request) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	fmt.Println(s.RoomList)

	if len(s.RoomList) == 0 {
		return
	}

	json.NewEncoder(w).Encode(s.RoomList)
}
