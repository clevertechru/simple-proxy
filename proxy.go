package main

import (
	"fmt"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"time"

	"github.com/clevertechru/simple-proxy/config"
)

var (
	Buildstamp string
	Commit     string
)
var startupTime = time.Now().Unix()

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("%s\n", resBody)
	}))

	config, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	var targets []*middleware.ProxyTarget
	for _, upstream := range config.Upstream {
		url, err := url.Parse(upstream)
		if err != nil {
			e.Logger.Fatal(err)
		}
		targets = append(targets, &middleware.ProxyTarget{URL: url})
	}

	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServicePort)))
}
