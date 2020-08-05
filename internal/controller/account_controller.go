package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-clean-architecture/internal/core/account"
	"net/http"
)

/*
Take a look here :)
1.- This file implements ServiceInterface methods. You can implement interface methods in separated files too.
2.- Errors raised at this layer should contain human readable messages to understand what went wrong
*/

func (w *GoClean) ListAccounts(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "ok")
}

func (w *GoClean) SaveAccount(ctx echo.Context) error {
	parsedAccount := account.Account{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&parsedAccount)
	if err != nil {
		logrus.Error("Return error message")
	}
	w.AccountService.SaveAccount(parsedAccount)
	return ctx.String(http.StatusOK, "ok")
}
