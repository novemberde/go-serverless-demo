package api

import (
	db "go-serverless-demo/internal/db"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	e.GET("/:username", a.getTodos)
	e.POST("/:username", a.newTodo)
	e.PUT("/:username/:created_at", a.updateTodo)
	e.DELETE("/:username/:created_at", a.deleteTodo)
	e.POST("/user", a.newUser)

	return a
}

func greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to go-serverless-demo! Please visit https://novemberde.github.io")
}
