package health_go

import (
	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v4"
)

type Service interface {
	HealthCheck() gin.HandlerFunc
}

type service struct{}

func New() Service {
	return &service{}
}

func (s service) HealthCheck() gin.HandlerFunc {
	h, _ := health.New()
	return gin.WrapF(h.HandlerFunc)
}