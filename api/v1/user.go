package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profile/auth"
	"profile/core"
	"time"
)

func (h *ViewHandler) Login(c echo.Context) error {
	data := c.(*core.CustomContext).GetBody()
	username := data["username"]
	password := data["password"]

	if username != "admin" || password != "123123" {
		return echo.ErrUnauthorized
	}

	t, err := auth.GenJWT("admin", true, []byte("secret-super-passwd"), time.Hour*6)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (h *ViewHandler) UserInfo(c echo.Context) error {
	claims := c.(*core.CustomContext).GetUser()
	return c.JSON(http.StatusOK, claims)
}

func (h *ViewHandler) Logout(c echo.Context) error {
	println("user logout")
	return c.JSON(http.StatusOK, echo.Map{})
}
