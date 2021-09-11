package d2k

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/migueleliasweb/d2k/src/openapigen"
)

type Handlers struct {
	openapigen.ServerInterface
}

// Ping
// (GET /_ping)
func (d2k *Handlers) SystemPing(ctx echo.Context) error {
	ctx.String(http.StatusOK, "OK")
	return nil
}

// Ping
// (HEAD /_ping)
func (d2k *Handlers) SystemPingHead(ctx echo.Context) error {
	ctx.String(http.StatusOK, "OK")
	return nil
}

// Create a container
// (POST /containers/create)
func (d2k *Handlers) ContainerCreate(
	ctx echo.Context,
	params openapigen.ContainerCreateParams,
) error {
	fmt.Println(params)
	return nil
}
