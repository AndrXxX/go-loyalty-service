package interfaces

import (
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type OrderService interface {
	Find(m *ormmodels.Order) *ormmodels.Order
	FindAll(m *ormmodels.Order) []*ormmodels.Order
	Create(m *ormmodels.Order) (*ormmodels.Order, error)
	Update(m *ormmodels.Order) error
}
