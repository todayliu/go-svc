package micro

import (
	"github.com/todayliu/go-svc/svc"
)

type IService interface {
	Init() error
	Start() error
	Stop(svc.Signal) error
}
