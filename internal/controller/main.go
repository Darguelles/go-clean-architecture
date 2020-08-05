package controller

import (
	"go-clean-architecture/internal/controller/openapi"
	"go-clean-architecture/internal/core/account"
	"github.com/labstack/echo/v4"
)

// GoClean contains the API implementation
type GoClean struct {
	Handler        openapi.ServerInterface
	AccountService account.AccountService
}

// NewAPI returns a new api instance
func NewAPI() *GoClean {
	service := account.AccountServiceInstance()
	return &GoClean{
		AccountService: service,
	}
}

// RegisterHandlers register handlers to echo server
func (w *GoClean) RegisterHandlers(e *echo.Echo) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		e.Logger.Fatal("Cound not get spec")
	}
	swagger.Servers = nil
	openapi.RegisterHandlers(e, w)
}
