package handlers

import (
  "net/http"
  "log"
  "io/ioutil"
  "fmt"
)

type Hello struct {
  logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
  return &Hello{ logger }
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  h.logger.Println("Hello World")
  // can read from r.Body from the user request: ioReader
  data, err := ioutil.ReadAll(r.Body)
  if err != nil {
    // w.WriteHeader(http.StatusBadRequest)
    // w.Write([]byte("Error!"))
    // equal:
    http.Error(w, "Oops", http.StatusBadRequest)
    return
  }
  h.logger.Printf("Data: %s\n", data)

  // write response to the user
  fmt.Fprintf(w, "Hello %s\n", data)
}
