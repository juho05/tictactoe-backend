package server

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	waitingClient *Client
	matches       []*Match
	matchesLock   sync.RWMutex
}

func New() *Server {
	return &Server{
		waitingClient: nil,
		matches:       make([]*Match, 1),
	}
}

func (s *Server) RemoveMatch(match *Match) {
	s.matchesLock.Lock()
	for i, m := range s.matches {
		if m == match {
			s.matches = append(s.matches[:i], s.matches[i+1:]...)
		}
	}
	s.matchesLock.Unlock()
}

func (s *Server) Listen(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	defer listener.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening on port %d...\n", port)

	for {
		con, err := listener.Accept()
		if err != nil {
			break
		}
		client := NewClient(s, con)
		go client.handleConnection()

		if s.waitingClient != nil {
			if s.waitingClient.send("ping") != nil {
				s.waitingClient = nil
			} else {
				match := s.NewMatch(s.waitingClient, client)
				s.matchesLock.Lock()
				s.matches = append(s.matches)
				s.matchesLock.Unlock()
				s.waitingClient = nil
				match.begin()
			}
		} else {
			s.waitingClient = client
		}
	}
}
