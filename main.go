package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
  // register handler for a path for default ServeMux
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    log.Println("Hello World")
    // can read from r.Body from the user request: ioReader
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
      // w.WriteHeader(http.StatusBadRequest)
      // w.Write([]byte("Error!"))
      // equal:
      http.Error(w, "Oops", http.StatusBadRequest)
      return
    }
    log.Printf("Data: %s\n", data)

    // write response to the user
    fmt.Fprintf(w, "Hello %s\n", data)
  })

  http.HandleFunc("/goodbye", func (http.ResponseWriter, *http.Request) {
    log.Println("Bye World")
  })

  // 2 arguments: address string, handler
  // can be '127.0.0.1:9090'
  http.ListenAndServe(":9090", nil)
}




// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// have lots of methods

// http.ServeMux: is a struct, responsible for redirecting paths, 
// you map a func to a path, and then ServeMux will decide which func gets executed
// DefaultServeMux: used by HandleFunc
// ServeMux: receive pattern which is a path, and a Handler

// HandleFunc: register the handler function for the given pattern in DefaultServeMux

// Handler: an interface of the ServeHTTP(ResponseWriter, *Request) method
