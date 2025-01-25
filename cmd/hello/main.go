package main

import (
	"hellogofr/internal/hello/server"

	_ "go.uber.org/automaxprocs"
)

func main() {
	server.New().Run()
}
