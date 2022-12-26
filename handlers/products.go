package handlers

import (
	// "encoding/json"
	"log"
	"net/http"

	"microservice/data"
)

type Products struct {
  logger *log.Logger
}

func NewProducts (logger *log.Logger) *Products {
  return &Products{logger}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
    p.getProducts(w, r)
    return
  }

  // catch all
  w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts (w http.ResponseWriter, r *http.Request) {
  lp := data.GetProducts()

  // return data as json
  // data, err := json.Marshal(lp)

  // use encoder for better performance
  err := lp.ToJSON(w)

  if err != nil {
    http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
  }

  // use this if using Marshal
  // w.Write(data)

}
