package jobs

import (
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type accrualClient interface {
	Fetch(order string) (statusCode int, m *entities.Accrual)
}

type updater interface {
	Update(m *ormmodels.Order) error
}
