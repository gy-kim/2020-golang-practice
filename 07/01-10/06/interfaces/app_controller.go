package interfaces

import (
	"log"

	"github.com/gy-kim/2020-golang-practice/07/01-10/06/infrastructure"
)

type AppController struct {
	productAPI infrastructure.ProductAPI
}

func (ac *AppController) New(productAPI infrastructure.ProductAPI) {
	ac.productAPI = productAPI
}

func (ac *AppController) PrintProducts() {
	productList, err := ac.productAPI.ProductList()
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range productList {
		log.Print(product.String())
	}
}
