package services

import (
	"context"
	"testing"

	"github.com/Bad4pple/Standardize/systems/tavern/domain/product"
	"github.com/Bad4pple/Standardize/systems/tavern/services/order"
	"github.com/google/uuid"
)

func TestTavern(t *testing.T) {
	products := init_products(t)
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://root:example@localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Starlette")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}
}
func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healty", 19)
	if err != nil {
		t.Fatal(err)
	}
	peanuts, err := product.NewProduct("Peanuts", "Snack", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	wine, err := product.NewProduct("Wine", "Healty", 99)
	if err != nil {
		t.Fatal(err)
	}
	return []product.Product{beer, peanuts, wine}
}
