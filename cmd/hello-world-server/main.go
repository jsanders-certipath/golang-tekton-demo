package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jsanders-certipath/tekton-demo/pkg/http/rest"
)

func main() {
	router := rest.Handler()
	figure.NewColorFigure("HELLO WORLD", "shadow", "green", true).Print()

	e := router.Run()
	if e != nil {
		panic(e)
	}
}
