package web

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

const (
	headerContentType = "Content-Type"

	mimeJSON = "application/json"
)

type Proccessor interface {
	Proccess(io.Reader, io.Writer) error
}

type Server struct {
	p Proccessor
}

func New(p Proccessor) *Server {
	return &Server{
		p: p,
	}
}

func (s *Server) Run(addr string) error {
	http.HandleFunc("/storage", s.Storage)

	return http.ListenAndServe(addr, nil)
}

func (s *Server) Storage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	buf := &bytes.Buffer{}
	err := s.p.Proccess(r.Body, buf)
	if err != nil {
		log.Println("Storage error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set(headerContentType, mimeJSON)
	w.WriteHeader(http.StatusOK)

	io.Copy(w, buf)
}
