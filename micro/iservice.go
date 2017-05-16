package micro

import (
	"github.com/todayliu/go-svc/svc"
)

type IService interface {
	Init(svc.Environment) error
	Start() error
	Stop(svc.Signal) error
}
