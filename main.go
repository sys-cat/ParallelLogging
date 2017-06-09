package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var s *echo.Echo

func main() {
	var std_port = flag.String("port", "8080", "Set Port")
	flag.Parse()
	port := ":" + *std_port
	s = echo.New()

	// Logger file
	file, err := os.OpenFile(fmt.Sprintf("logger.%s.log", time.Now().Format("20060102")), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	s.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: file,
	}))
	s.Use(middleware.Recover())
	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "https://localhost"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	s.GET("/", getter)

	listen, err := net.Listen("tcp4", port)
	if err != nil {
		s.Logger.Fatal("Cant Listen tcp4")
	}
	s.Listener = listen
	s.Logger.Fatal(s.Server.Serve(s.Listener))
}

func getter(c echo.Context) error {
	s.Logger.Info("hello")
	return c.JSON(http.StatusOK, "Hello")
}
