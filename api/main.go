package main

import (
	"os"
	"strconv"

	ServerFactory "github.com/vmustillo/vue-go/api/server"
)

var port int

func init() {
	rawPort := os.Getenv("PORT")

	if len(rawPort) > 0 {
		var err error
		port, err = strconv.Atoi(rawPort)
		if err != nil {
			panic(err)
		}
	} else {
		port = 8000
	}
}

func main() {
	server := ServerFactory.NewServer(port)

	server.Start()
}