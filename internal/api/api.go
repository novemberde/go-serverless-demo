package api

import (
	db "go-serverless-demo/internal/db"
	"go-serverless-demo/internal/middlewares"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// API hold echo instance
type API struct {
	*echo.Echo
	db *db.DB
}

// New returns echo
func New() *API {
	e := echo.New()
	a := &API{
		Echo: e,
		db:   db.New(os.Getenv("DYNAMO_REGION"), os.Getenv("DYNAMO_TABLE_NAME")),
	}

	e.Use(middlewares.CORS)
	e.GET("/", greeting)
	e.GET("/:username", a.getTodos)
	e.POST("/:username", a.newTodo)
	e.PUT("/:username/:created_at", a.updateTodo)
	e.DELETE("/:username/:created_at", a.deleteTodo)
	e.POST("/user", func(c echo.Context) error {
		un := c.Param("username")

		if err := a.db.AddUser(un); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "ok")
	})

	return a
}

func greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to go-serverless-demo! Please visit https://novemberde.github.io")
}
