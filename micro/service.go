package micro

import (
	"log"
	"os"

	"github.com/todayliu/go-svc/svc"
)

type Service struct {
	logFile *os.File
	Runner  IService
}

func NewServer(isv IService) *Service {
	return &Service{Runner: isv}
}

func (s *Service) Init(env svc.Environment) error {
	 return s.Runner.Init(env)
}

func (s *Service) Start() error {
	if s.Runner != nil {
		go s.Runner.Start()
	}

	return nil
}

func (s *Service) Stop(signal svc.Signal) error {
	log.Printf("out %s", signal.String())
	if s.Runner != nil {
		s.Runner.Stop(signal)
	}
	return nil
}
