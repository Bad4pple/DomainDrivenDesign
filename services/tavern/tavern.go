package services

import (
	"log"

	"github.com/Bad4pple/Standardize/systems/tavern/services/order"
	"github.com/google/uuid"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService *order.OrderService
	// Billing Service
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}
	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer_id uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer_id, products)
	if err != nil {
		return err
	}
	log.Printf("\nBill the customer: %0.0f", price)
	return nil
}
