package main

import (
	"log"

	"flamingo.me/dingo"
	"github.com/gy-kim/2020-golang-practice/07/01-10/06/interfaces"
)

func main() {
	// productAPI := infrastructure.ConcreteProductAPI{ApiUrl: ProductAPI}

	// appController := interfaces.AppController{}
	// appController.New(&productAPI)

	// appController.PrintProducts()

	injector, err := dingo.NewInjector(new(Module))
	if err != nil {
		panic(err)
	}

	instance, err := injector.GetInstance(new(interfaces.AppController))
	if err != nil {
		log.Fatal(err)
	}

	appController := instance.(*interfaces.AppController)
	appController.PrintProducts()
}
