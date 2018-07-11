package main

import (
	"log"

	"github.com/fschnko/kvs/server"
	"github.com/fschnko/kvs/tcp"
	"github.com/fschnko/kvs/web"
)

func main() {
	s := server.New()

	go func() { log.Fatal(tcp.New(s).Run("localhost:8081")) }()

	log.Fatal(web.New(s).Run("localhost:8080"))

}
