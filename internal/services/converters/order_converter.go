package converters

import (
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type orderConverter struct {
}

func NewOrderConverter() *orderConverter {
	return &orderConverter{}
}

func (c orderConverter) Convert(m *ormmodels.Order) *entities.Order {
	return &entities.Order{
		Number:    m.Number,
		Status:    m.Status,
		Accrual:   m.Accrual,
		CreatedAt: entities.RFC3339Time{Time: m.CreatedAt},
	}
}

func (c orderConverter) ConvertMany(list []*ormmodels.Order) []*entities.Order {
	result := make([]*entities.Order, 0, len(list))
	for _, order := range list {
		result = append(result, c.Convert(order))
	}
	return result
}
