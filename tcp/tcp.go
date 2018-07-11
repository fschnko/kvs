package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Proccesor interface {
	Proccess(io.Reader, io.Writer) error
}

type Server struct {
	p Proccesor
}

func New(p Proccesor) *Server {
	return &Server{
		p: p,
	}
}

func (s *Server) Run(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("new listener: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept tcp connection: %v \n", err)
			continue
		}
		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()
	err := s.p.Proccess(conn, conn)
	if err != nil {
		log.Printf("proccess tcp connection: %v \n", err)
	}
}
