package api

import (
	db "go-serverless-demo/internal/db"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New Return echo
func New() *echo.Echo {
	e := echo.New()
	d := db.New(os.Getenv("DYNAMO_REGION"), os.Getenv("DYNAMO_TABLE_NAME"))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5000", "http://novemberde-go-todo.s3-website.ap-northeast-2.amazonaws.com/", "https://d1ek9bfmwa0wns.cloudfront.net/"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to go-serverless-demo! Please visit https://novemberde.github.io")
	})
	e.GET("/:username", func(c echo.Context) error {
		un := c.Param("username")
		todos, err := d.Find(un)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, todos)
	})
	e.POST("/:username", func(c echo.Context) error {
		un := c.Param("username")
		t := new(db.Todo)
		if err := c.Bind(t); err != nil {
			return err
		}
		t.Username = un
		t.UserAgent = c.Request().Header.Get("User-Agent")
		err := d.Create(t)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, "created")
	})
	e.PUT("/:username/:created_at", func(c echo.Context) error {
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
		err = d.Update(t)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusNoContent, "updated")
	})
	e.DELETE("/:username/:created_at", func(c echo.Context) error {
		un := c.Param("username")
		ca := c.Param("created_at")

		tt, err := time.Parse(time.RFC3339, ca)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if err := d.Delete(&db.Todo{
			Username:  un,
			CreatedAt: tt,
		}); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusNoContent, "deleted")
	})

	e.POST("/user", func(c echo.Context) error {
		un := c.Param("username")

		if err := d.AddUser(un); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "ok")
	})

	return e
}
