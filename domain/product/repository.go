package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("no such product")
	ErrProductAlreadyExits = errors.New("there is already such an product")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(uuid.UUID) (Product, error)
	Add(Product) error
	Update(Product) error
	Delete(uuid.UUID) error
}
