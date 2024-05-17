package main

import (
	"cop/internal"
	"net/http"

	_ "github.com/leapkit/core/envload"
	"github.com/leapkit/core/server"
)

func main() {
	r := server.New()

	internal.SetRoutes(r)

	err := http.ListenAndServe(r.Addr(), r.Handler())
	if err != nil {
		panic(err)
	}
}
