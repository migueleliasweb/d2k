package main

import (
	"github.com/labstack/echo/v4"
	"github.com/migueleliasweb/docker-2-k8s/src/openapigen"
)

type docker2K8sHandlers struct {
	openapigen.ServerInterface
}

func main() {
	handlers := &docker2K8sHandlers{}
	e := echo.New()
	openapigen.RegisterHandlers(e, handlers)
}
