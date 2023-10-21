package products

import (
	"encoding/json"
	"estore-gin/commons"
	"os"
	"strconv"
)

const dataPath = "products/products.json"

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductRepository struct{}

func (p *ProductRepository) FindById(id string) (Product, error) {
	var products []Product
	var result Product
	err := commons.ReadJSON(dataPath, &products)
	if err != nil {
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
	if err != nil {
		return nil, err
	}

	// TODO: filter by query
	return products, nil
}

func (p *ProductRepository) Create(product Product) (Product, error) {
	var result Product
	file, err := os.OpenFile(dataPath, os.O_RDWR|os.O_CREATE, os.ModeExclusive)
	defer file.Close()

	if err != nil {
		return result, err
	}

	// if err = unix.Flock(int(file.Fd()), unix.LOCK_EX); err != nil {
	// 	return result, err
	// }
	// defer unix.Flock(int(file.Fd()), unix.LOCK_UN)

	var data []Product
	json.NewDecoder(file).Decode(&data)

	id := 1
	for _, v := range data {
		_id, err := strconv.Atoi(v.ID)
		if err != nil {
			return result, err
		}
		if id < _id {
			id = _id + 1
		}
	}
	product.ID = strconv.Itoa(id)
	data = append(data, product)

	buffer, convertErr := json.MarshalIndent(data, "", "  ")
	if convertErr != nil {
		return result, convertErr
	}

	if _, err = file.WriteAt(buffer, 0); err != nil {
		return result, err
	}

	return product, nil
}