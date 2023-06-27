package memory

import (
	"sync"

	"github.com/Bad4pple/Standardize/systems/tavern/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(cus customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customers[cus.GetID()]; ok {
		return customer.ErrFailedToAddCustomer
	}

	mr.Lock()
	mr.customers[cus.GetID()] = cus
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(cus customer.Customer) error {
	if _, ok := mr.customers[cus.GetID()]; !ok {
		return customer.ErrCustomerNotFound
	}
	mr.Lock()
	mr.customers[cus.GetID()] = cus
	mr.Unlock()

	return nil
}
