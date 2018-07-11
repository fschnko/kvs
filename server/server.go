package server

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/fschnko/kvs/storage"
)

type Server struct {
	s storage.Storage
}

func New() *Server {
	return &Server{
		s: storage.New(),
	}
}

func (s *Server) Proccess(in io.Reader, out io.Writer) error {
	req, err := read(in)
	if err != nil {
		return err
	}

	resp := &Response{
		Method: req.Method,
		Key:    req.Key,
	}

	switch req.Method {
	case Get:
		s.get(req, resp)
	case Set:
		s.set(req, resp)
	case Delete:
		s.delete(req, resp)
	case Exists:
		s.exists(req, resp)
	default:
		return fmt.Errorf("unsupported method: %v", req.Method)
	}

	return write(out, resp)
}

func (s *Server) get(req *Request, resp *Response) {
	val, err := s.s.Get(req.Key)
	if err != nil {
		resp.Error = err
		return
	}
	resp.Value = val
}

func (s *Server) set(req *Request, resp *Response) {
	err := s.s.Set(req.Key, req.Value)
	if err != nil {
		resp.Error = err
		return
	}
	resp.Value = req.Value
	resp.Result = SuccessResult
}

func (s *Server) delete(req *Request, resp *Response) {
	err := s.s.Delete(req.Key)
	if err != nil {
		resp.Error = err
		return
	}

	resp.Result = SuccessResult
}

func (s *Server) exists(req *Request, resp *Response) {
	exist, err := s.s.Exists(req.Key)
	if err != nil {
		resp.Error = err
		return
	}

	if exist {
		resp.Result = ExistResult
		return
	}
	resp.Result = NotExist
}

func read(r io.Reader) (*Request, error) {
	req := new(Request)
	if err := json.NewDecoder(r).Decode(req); err != nil {
		return nil, fmt.Errorf("decode: %v", err)
	}
	return req, nil
}

func write(w io.Writer, resp *Response) error {
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return fmt.Errorf("encode: %s", err)
	}

	return nil
}
