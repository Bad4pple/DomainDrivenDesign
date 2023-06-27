package memory

import (
	"errors"
	"testing"

	"github.com/Bad4pple/Standardize/systems/tavern/domain/customer"
	"github.com/google/uuid"
)

func TestMemoryGetCustomer(t *testing.T) {
	type testCase struct {
		id          uuid.UUID
		name        string
		exceptedErr error
	}
	cust, err := customer.NewCustomer("saharat")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}
	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("ad558232-cc1b-4986-b0b5-3b4f4b9f313f"),
			exceptedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			exceptedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(tc.exceptedErr, err) {
				t.Error(customer.ErrCustomerNotFound)
			}
		})
	}
}
