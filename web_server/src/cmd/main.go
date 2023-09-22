package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mono424/go-pts"
	ptsc_gorilla "github.com/mono424/go-pts-gorilla-connector"
	"github.com/tarm/serial"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
)

var appPath = "/Users/khadim/dev/tisch/web_server/app"

var config = serial.Config{
	Name: "/dev/ttyS0",
	Baud: 9600,
}

func main() {
	r := gin.Default()
	tubeSystem := pts.New(ptsc_gorilla.NewConnector(
		websocket.Upgrader{},
		func(err *pts.Error) {
			println(err.Description)
		},
	))

	tubeSystem.RegisterChannel("/control", pts.ChannelHandlers{
		OnSubscribe: func(s *pts.Context) {
			println("Client joined: " + s.FullPath)
		},
		OnMessage: func(s *pts.Context, message *pts.Message) {
			println("New Message on " + s.FullPath + ": " + string(message.Payload))
		},
		OnUnsubscribe: func(s *pts.Context) {
			println("Client left: " + s.FullPath)
		},
	})

	r.GET("/connect", func(c *gin.Context) {
		properties := make(map[string]interface{}, 1)
		properties["ctx"] = c

		if err := tubeSystem.HandleRequest(c.Writer, c.Request, properties); err != nil {
			println("Something went wrong while handling a Socket request")
		}
	})
	r.NoRoute(ReverseProxy)

	go func() {
		cmd := exec.Command("yarn", "dev")
		cmd.Dir = appPath
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	r.Run()
}

func ReverseProxy(c *gin.Context) {
	remote, _ := url.Parse("http://localhost:3000")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL = c.Request.URL
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
