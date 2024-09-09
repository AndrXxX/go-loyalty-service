package interfaces

import (
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type WithdrawService interface {
	Find(m *ormmodels.Withdraw) *ormmodels.Withdraw
	FindAll(m *ormmodels.Withdraw) []*ormmodels.Withdraw
	Create(m *ormmodels.Withdraw) (*ormmodels.Withdraw, error)
}
