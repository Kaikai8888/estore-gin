package products

import (
	"estore-gin/commons"
)

const dataPath = "products/products.json"


type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

type ProductRepository struct {}

func (p *ProductRepository) FindById(id string) (Product, error) {
  var products []Product
  var result Product
	err := commons.ReadJSON(dataPath, &products)
  if err!= nil {
    return result, err
  }

  // TODO: write array methods into commons package for reuse
  
  for i := 0; i < len(products); i++ {
    if products[i].ID == id {
      return products[i], nil
    }
  }
  return result, nil
}

// query Product
func (p *ProductRepository) Find() ([]Product, error) {
  products := []Product{}
	err := commons.ReadJSON(dataPath, &products)
  if err!= nil {
    return nil, err
  }

  // TODO: filter by query
  return products, nil
}









// var products []Product
// err = json.Unmarshal(jsonFile, &products)
// if err!= nil {
//   return nil, err
// }

// for _, product := range products {
//   if product.ID == id {
//     return &product, nil
//   }
// }

// return nil, nil
