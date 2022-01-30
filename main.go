package main

import (
	echoPrometheus "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

func main() {
	requests := promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_request_total",
		Help: "The total number of api called",
	})

	e := echo.New()
	p := echoPrometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.GET("/hello", func(c echo.Context) error {
		requests.Inc()
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Logger.Fatal(e.Start(":8081"))
}