package main

import (
	"context"
	"log"

	"github.com/Bad4pple/Standardize/systems/tavern/domain/product"
	"github.com/Bad4pple/Standardize/systems/tavern/services/order"
	tavern "github.com/Bad4pple/Standardize/systems/tavern/services/tavern"
	"github.com/google/uuid"
)

func main() {
	products := productInventory()
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://root:example@localhost:27017/"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		log.Println(err)
	}
	tavern_service, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)

	if err != nil {
		log.Println(err)
	}

	uid, err := os.AddCustomer("Standardize")
	if err != nil {
		log.Println(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern_service.Order(uid, order)
	if err != nil {
		log.Println(err)
	}

}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healty", 19)
	if err != nil {
		panic(err)
	}
	peanuts, err := product.NewProduct("Peanuts", "Snack", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healty", 99)
	if err != nil {
		panic(err)
	}
	return []product.Product{beer, peanuts, wine}
}
