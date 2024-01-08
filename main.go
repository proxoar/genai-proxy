package main

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	mapping := map[string]TargetUrl{
		"gpt-proxy":    {"https", "api.openai.com", ""},
		"gemini-proxy": {"https", "generativelanguage.googleapis.com", ""},
	}
	e := echo.New()
	e.HideBanner = true
	e.Debug = true
	e.Use(echoprometheus.NewMiddleware("proxy_over_proxy"))
	e.GET("/metrics", echoprometheus.NewHandler())
	e.GET("/", OkAlive)
	e.OPTIONS("/*", Options)
	e.Any("/*", ProxyOverProxy(mapping))
	err := e.Start(":8000")
	if err != nil {
		panic(err)
	}
}
