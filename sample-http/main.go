package main

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo"
)

type Enforcer struct {
	enforcer *casbin.Enforcer
}

func (e *Enforcer) Enforce(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _, _ := c.Request().BasicAuth()
		method := c.Request().Method
		path := c.Request().URL.Path

		result, _ := e.enforcer.Enforce(user, path, method)

		if result {
			return next(c)
		}
		return echo.ErrForbidden
	}
}

func main() {
	e := echo.New()
	newEnforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		return
	}
	enforcer := Enforcer{
		enforcer: newEnforcer,
	}

	e.Use(enforcer.Enforce)

	e.GET("/project", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "project get allowed")
	})
	e.POST("/project", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "project post allowed")
	})

	e.GET("/channel", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "channel get allowed")
	})

	e.POST("/channel", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "channel post allowed")
	})
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
