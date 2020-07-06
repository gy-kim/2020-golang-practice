package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gy-kim/2020-golang-practice/07/01-10/06/domain"
)

type ProductAPI interface {
	ProductList() (domain.ProductList, error)
}

type ConcreteProductAPI struct {
	ApiUrl string `inject:"config:productApiUrl"`
}

func (cpa *ConcreteProductAPI) Get(endpoint string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", cpa.ApiUrl, endpoint))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (cpa *ConcreteProductAPI) ProductList() (domain.ProductList, error) {
	var err error
	var productList domain.ProductList
	body, err := cpa.Get("/products")
	if err != nil {
		return productList, err
	}

	err = json.Unmarshal(body, &productList)
	if err != nil {
		return productList, err
	}

	return productList, nil
}
