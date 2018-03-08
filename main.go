package main

import (
	"github.com/rusucosmin/goworkshop/persistence"
	"github.com/rusucosmin/goworkshop/web"
)

func main() {
	if err := persistence.InitDB(); err != nil {
		panic(err)
	}

	server := web.RestServer{
		Port: 8080,
	}

	server.StartServer()
}
