package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
  // - is used to ignore the field
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Marshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
  encoder := json.NewEncoder(w)
  return encoder.Encode(p)
}
func (p *Product) FromJSON(r io.Reader) error {
  decoder := json.NewDecoder(r)
  return decoder.Decode(p)
}

// Put functions interacting with DB here
func GetProducts() Products {
  return productList
}
func AddProduct(p *Product) {
  p.ID = getNextId()
  productList = append(productList, p)
}
func UpdateProduct(id int, p *Product) error {
  _, pos, err := findProduct(id)  
  if err != nil {
    return err
  }

  p.ID = id
  productList[pos] = p
  return nil
}

// helpers
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
  for i, p := range productList {
    if id == p.ID {
      return p, i, nil
    }
  }

  return nil, -1, ErrProductNotFound
}
func getNextId() int {
  return len(productList) + 1
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
