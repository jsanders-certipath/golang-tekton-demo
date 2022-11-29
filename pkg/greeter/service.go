package greeter

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	SayHello(ctx *gin.Context) string
}

type service struct{}

func New() Service {
	return &service{}
}

func (s service) SayHello(ctx *gin.Context) string {

	return "Hello world!"
}
