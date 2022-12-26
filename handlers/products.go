package handlers

import (
	// "encoding/json"
	"log"
	"net/http"
	"strconv"

	"microservice/data"

	"github.com/gorilla/mux"
)

type Products struct {
  logger *log.Logger
}

func NewProducts (logger *log.Logger) *Products {
  return &Products{logger}
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
  // To get the variables from the URL
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    http.Error(w, "Converting error", http.StatusBadRequest)
  }

  prod := &data.Product{}
  
  err = prod.FromJSON(r.Body)
  if err != nil {
    http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
  }

  err = data.UpdateProduct(id, prod)
  if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
  if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
  }
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
  // create new product object
  prod := &data.Product{}

  // get data from request
  err := prod.FromJSON(r.Body)
  if err != nil {
    http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
  }

  p.logger.Printf("Prod: %#v\n", prod)

  // add to array
  data.AddProduct(prod) 
}

func (p *Products) GetProducts (w http.ResponseWriter, r *http.Request) {
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
