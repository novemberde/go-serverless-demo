package middlewares

import (
	"net/http"

	"github.com/labstack/echo/middleware"
)

var allowOrigins = []string{
	"http://localhost:5000",
	"https://go-todo.judoka.dev",
	"http://novemberde-go-todo.s3-website.ap-northeast-2.amazonaws.com",
	"https://d1ek9bfmwa0wns.cloudfront.net",
}

var allowMethods = []string{
	http.MethodGet,
	http.MethodPut,
	http.MethodPost,
	http.MethodDelete,
	http.MethodOptions,
}

// CORS is CORS config options using Echo framework.
var CORS = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: allowOrigins,
	AllowMethods: allowMethods,
})
