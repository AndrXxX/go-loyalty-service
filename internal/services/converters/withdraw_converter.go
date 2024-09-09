package converters

import (
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type withdrawConverter struct {
}

func NewWithdrawConverter() *withdrawConverter {
	return &withdrawConverter{}
}

func (c withdrawConverter) Convert(m *ormmodels.Withdraw) *entities.Withdraw {
	return &entities.Withdraw{
		Order:     m.Order,
		Sum:       m.Sum,
		CreatedAt: entities.RFC3339Time{Time: m.CreatedAt},
	}
}

func (c withdrawConverter) ConvertMany(list []*ormmodels.Withdraw) []*entities.Withdraw {
	result := make([]*entities.Withdraw, 0, len(list))
	for _, order := range list {
		result = append(result, c.Convert(order))
	}
	return result
}
