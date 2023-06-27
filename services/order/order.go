package order

import (
	"context"
	"log"

	"github.com/Bad4pple/Standardize/systems/tavern/domain/customer"
	"github.com/Bad4pple/Standardize/systems/tavern/domain/customer/memory"
	mongo_db "github.com/Bad4pple/Standardize/systems/tavern/domain/customer/mongo"
	"github.com/Bad4pple/Standardize/systems/tavern/domain/product"
	prodmem "github.com/Bad4pple/Standardize/systems/tavern/domain/product/memory"
	"github.com/google/uuid"
)

type OrderService struct {
	customers customer.CustomerRespository
	products  product.ProductRepository
	// billing billing.Service
}

type OrderConfiguration func(os *OrderService) error

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRespository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo_db.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {

	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []product.Product
	var total float64
	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has orderd %d products", c.GetID(), len(products))

	return total, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	// cust
	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}
	return c.GetID(), nil
}
