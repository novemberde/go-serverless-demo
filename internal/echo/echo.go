package echo

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"go-serverless-demo/internal/api"
)

func NewEcho(a *api.API) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5000",
			"https://go-todo.judoka.dev",
			"http://novemberde-go-todo.s3-website.ap-northeast-2.amazonaws.com",
			"https://d1ek9bfmwa0wns.cloudfront.net"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions},
	}))

	e.GET("/", greeting)

	e.GET("/:username", a.GetTodos)
	e.POST("/:username", a.NewTodo)
	e.PUT("/:username/:created_at", a.UpdateTodo)
	e.DELETE("/:username/:created_at", a.DeleteTodo)

	return e
}

func greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to go-serverless-demo! Please visit https://novemberde.github.io")
}
