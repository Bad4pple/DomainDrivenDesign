package customer

import (
	"errors"

	"github.com/Bad4pple/Standardize/systems/tavern"
	"github.com/google/uuid"
)

var ErrInvalidPerson = errors.New("person's name has to have valid name")

type Customer struct {
	person      *tavern.Person
	product     []*tavern.Item
	transaction []tavern.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:      person,
		product:     make([]*tavern.Item, 0),
		transaction: make([]tavern.Transaction, 0),
	}, nil
}

func (c Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}

func (c Customer) GetName() string {
	return c.person.Name
}
