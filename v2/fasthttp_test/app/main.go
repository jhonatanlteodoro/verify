package main

import (
	"flag"

	initalization "github.com/jhonatanlteodoro/fasthttp_test/app/initialization"
)

func main() {
	var port string
	var host string

	flag.StringVar(&port, "port", "8089", "Server port to listen on")
	flag.StringVar(&host, "host", "localhost", "App host")

	App := initalization.App{}
	App.Initilize()
	App.Run(host, port)
}
