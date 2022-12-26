package handlers

import (
	// "encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

  if r.Method == http.MethodPost {
    p.addProduct(w, r)
    return
  }

  if r.Method == http.MethodPut {
    // expect the id in the url
    regex := regexp.MustCompile(`/([0-9]+)`)
    path := r.URL.Path

    g := regex.FindAllStringSubmatch(path, -1)
    if len(g) != 1 {
      p.logger.Println("Invalid URI: more than one id")
      http.Error(w, "Invalid URI", http.StatusBadRequest)
      return
    }

    if len(g[0]) != 2 {
      p.logger.Println("Invalid URI: more than one captured group")
      http.Error(w, "Invalid URI", http.StatusBadRequest)
      return
    }

    idString := g[0][1]
    id, err := strconv.Atoi(idString)
    if err != nil {
      p.logger.Println("Invalid URI: converting error")
      http.Error(w, "Invalid URI", http.StatusBadRequest)
      return
    }

    p.updateProduct(id, w, r)
    return
  }

  // catch all
  w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
  prod := &data.Product{}
  
  err := prod.FromJSON(r.Body)
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

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
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
