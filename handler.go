package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httputil"
	"path"
)

type TargetUrl struct {
	Scheme     string
	Host       string
	PathPrefix string
}

func ProxyOverProxy(mapping map[string]TargetUrl) echo.HandlerFunc {
	return func(c echo.Context) error {
		rewriter := func(req *httputil.ProxyRequest) {
			originURL := req.Out.RequestURI
			originHost := req.Out.Host
			if t, ok := mapping[originHost]; ok {
				if t.Scheme != "" {
					req.Out.URL.Scheme = t.Scheme
				}
				if t.Host != "" {
					req.Out.Host = t.Host
					req.Out.URL.Host = t.Host
				}
				if t.PathPrefix != "" {
					req.Out.URL.Path = path.Join("/", t.PathPrefix, req.In.URL.Path)
				}
				log.Printf("proxying request %s -> %s\n", originURL, req.Out.URL.String())
			}
		}
		server := httputil.ReverseProxy{Rewrite: rewriter}
		server.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

func Options(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Oriecho", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	return c.String(http.StatusOK, "")
}

func OkAlive(c echo.Context) error {
	return c.String(http.StatusOK, "I'm alive.")
}
