package svc

import (
	"log"
	"os"
	"errors"
	"syscall"
)

type Server struct {
	logFile *os.File
	service  Service
}

func NewServer(isv Service) *Server {
	return &Server{service: isv}
}

func (s *Server) Init(env Environment) error {
	if s.service != nil {
		return s.service.Init(env)
	}
	return errors.New("服务对象不存在，Servers is not exist!")
}

func (s *Server) Start() error {
	if s.service != nil {
		go s.service.Start()
		return nil
	}
	return errors.New("服务对象不存在，Servers is not exist!")
}

func (s *Server) Stop(signal Signal) error {
	log.Printf("out %s", signal.String())
	if s.service != nil {
		s.service.Stop(signal)
	}
	return nil
}

func (s *Server) Watch(signals ...os.Signal)error{
	if len(signals) == 0{
		signals = append(signals,syscall.SIGINT, syscall.SIGTERM)
	}
	return Run(s,signals...);
}