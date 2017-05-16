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

func (this *Service) Init(env svc.Environment) error {

	return nil
}

func (this *Service) Start() error {

	if this.Runner != nil {
		go this.Runner.Start()
	}

	return nil
}

func (this *Service) Stop(signal svc.Signal) error {
	log.Printf("out %s", signal.String())
	if this.Runner != nil {
		this.Runner.Stop(signal)
	}
	return nil
}
