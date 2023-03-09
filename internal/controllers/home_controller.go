package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func Home(ctx echo.Context) error {
	fmt.Fprintf(ctx.Response(), "Welcome to GoShop Home Page")
	return nil
}
