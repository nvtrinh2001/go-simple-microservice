package main

import (
	"log"
	"net/http"
	"os"

  "./handlers"
)

func main() {
  l := log.New(os.Stdout, "simple-api", log.LstdFlags)
  hh := handlers.NewHello(l)

  sm := http.NewServeMux()
  sm.Handle("/", hh)

  // 2 arguments: address string, handler
  // can be '127.0.0.1:9090'
  http.ListenAndServe(":9090", sm)
}

