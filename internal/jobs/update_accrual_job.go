package jobs

import (
	"github.com/AndrXxX/go-loyalty-service/internal/enums/accrualstatuses"
	"github.com/AndrXxX/go-loyalty-service/internal/enums/orderstatuses"
	"github.com/AndrXxX/go-loyalty-service/internal/interfaces"
	"github.com/AndrXxX/go-loyalty-service/internal/ormmodels"
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const repeatTime = 1 * time.Second

type updateAccrualJob struct {
	ac accrualClient
	o  *ormmodels.Order
}

func NewUpdateAccrualJob(ac accrualClient, o *ormmodels.Order) interfaces.QueueJob {
	return &updateAccrualJob{ac, o}
}

func (j *updateAccrualJob) Execute() error {
	for {
		code, info := j.ac.Fetch(j.o.Number)
		if code != http.StatusOK {
			logger.Log.Info("got code from accrual", zap.Int("code", code))
			time.Sleep(repeatTime)
			continue
		}
		logger.Log.Info("got data from accrual", zap.Any("info", info))
		switch info.Status {
		case accrualstatuses.Invalid:
			j.o.Status = orderstatuses.Invalid
			// TODO: save model
			return nil
		case accrualstatuses.Registered:
			time.Sleep(repeatTime)
			continue
		case accrualstatuses.Processing:
			if j.o.Status != orderstatuses.Processing {
				j.o.Status = orderstatuses.Processing
				// TODO: save model
			}
			time.Sleep(repeatTime)
			continue
		case accrualstatuses.Processed:
			j.o.Status = orderstatuses.Processed
			j.o.Accrual = info.Accrual
			// TODO: save model
			return nil
		}
	}
}
