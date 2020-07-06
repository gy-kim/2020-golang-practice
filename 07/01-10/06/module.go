package main

import (
	"flamingo.me/dingo"
	"github.com/gy-kim/2020-golang-practice/07/01-10/06/infrastructure"
)

const ProductApiUrl = "https://my-product-api.com"

type Module struct{}

func (m *Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(string)).AnnotatedWith("config:proeuctApiUrl").ToInstance(ProductApiUrl)
	injector.Bind(new(infrastructure.ProductAPI)).To(infrastructure.ConcreteProductAPI{})
}
