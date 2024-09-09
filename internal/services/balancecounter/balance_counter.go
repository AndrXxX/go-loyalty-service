package balancecounter

import (
	"github.com/AndrXxX/go-loyalty-service/internal/entities"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type balanceCounter struct {
	ws interfaces.WithdrawService
	os interfaces.OrderService
}

func New(ws interfaces.WithdrawService, os interfaces.OrderService) *balanceCounter {
	return &balanceCounter{ws, os}
}

func (c *balanceCounter) Count(u *ormmodels.User) *entities.Balance {
	var b entities.Balance
	var oSum float64
	var wSum float64
	wList := c.ws.FindAll(&ormmodels.Withdraw{AuthorID: u.ID})
	for _, w := range wList {
		wSum += *w.Sum
	}
	oList := c.os.FindAll(&ormmodels.Order{AuthorID: u.ID})
	for _, o := range oList {
		if o.Accrual != nil {
			oSum += *o.Accrual
		}
	}
	current := oSum - wSum
	b.Current = &current
	b.Withdrawn = &wSum
	return &b
}
