package order

import (
	"testing"

	"github.com/Bad4pple/Standardize/systems/tavern/domain/customer"
	"github.com/Bad4pple/Standardize/systems/tavern/domain/product"
	"github.com/google/uuid"
)

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

func TestOrderNewOrderService(t *testing.T) {

	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Saharat")
	if err != nil {
		t.Fatal(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}
}
