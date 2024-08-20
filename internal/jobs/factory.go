package jobs

import (
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
)

type factory struct {
	ac accrualClient
	ou orderUpdater
}

func Factory(
	ac accrualClient,
	ou orderUpdater,
) *factory {
	return &factory{ac, ou}
}

func (f *factory) NewUpdateAccrualJob(o *ormmodels.Order) interfaces.QueueJob {
	return &updateAccrualJob{f.ac, o, f.ou}
}
