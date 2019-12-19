package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"go-serverless-demo/internal/db"
)

func (a *API) GetTodos(c echo.Context) error {
	un := c.Param("username")
	todos, err := a.db.Find(un)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func (a *API) NewTodo(c echo.Context) error {
	un := c.Param("username")
	t := new(db.Todo)
	if err := c.Bind(t); err != nil {
		return err
	}
	t.Username = un
	t.UserAgent = c.Request().Header.Get("User-Agent")
	err := a.db.Create(t)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "created")
}

func (a *API) UpdateTodo(c echo.Context) error {
	un := c.Param("username")
	ca := c.Param("created_at")
	t := new(db.Todo)
	if err := c.Bind(t); err != nil {
		return err
	}
	t.Username = un
	tt, err := time.Parse(time.RFC3339, ca)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	t.CreatedAt = tt
	err = a.db.Update(t)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, "updated")
}

func (a *API) DeleteTodo(c echo.Context) error {
	un := c.Param("username")
	ca := c.Param("created_at")

	tt, err := time.Parse(time.RFC3339, ca)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := a.db.Delete(&db.Todo{
		Username:  un,
		CreatedAt: tt,
	}); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusNoContent, "deleted")
}
