package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *API) newUser(c echo.Context) error {
	un := c.Param("username")

	if err := a.db.AddUser(un); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
