package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsanders-certipath/tekton-demo/pkg/greeter"
	healthgo "github.com/jsanders-certipath/tekton-demo/third_party/health-go"
	"github.com/jsanders-certipath/tekton-demo/third_party/ratelimiter"
)

var prefix = "/api/v1"

func Handler() *gin.Engine {

	router := gin.Default()

	//Establish healthcheck
	router.GET(prefix+"/health", healthgo.New().HealthCheck())

	//Setup security on all future endpoints
	router.Use(ratelimiter.New().InMemRateLimiter())

	//GET endpoints
	router.GET(prefix+"/hello", helloWorld())

	//POST endpoints
	return router
}

func helloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData := []byte(greeter.New().SayHello(c))
		c.Data(http.StatusOK, "application/json", jsonData)
		c.Next()
	}
}
