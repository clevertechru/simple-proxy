package main

import (
	"flag"
	"fmt"
	"net/http"

	"runtime"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Buildstamp string
	Commit     string
)
var startupTime = time.Now().Unix()

//go build && ./upstream -name echo-1 -port 3001 && ./upstream -name echo-2 -port 3002

func main() {
	name := flag.String("name", "echo", "server name")
	port := flag.String("port", "3000", "server port")
	flag.Parse()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("<div style='font-size: 8em;'>Hello from upstream server %s!</div>", *name))
	})
	e.GET("/alive", func(c echo.Context) error {
		data := map[string]interface{}{
			"alive":         true,
			"hostname":      "localhost:" + *port,
			"serviceName":   *name,
			"num_cpu":       runtime.NumCPU(),
			"num_goroutine": runtime.NumGoroutine(),
			"go_version":    runtime.Version(),
			"build_date":    Buildstamp,
			"commit":        Commit,
			"startup_time":  startupTime,
		}
		return c.JSON(http.StatusOK, data)
	})

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
}
