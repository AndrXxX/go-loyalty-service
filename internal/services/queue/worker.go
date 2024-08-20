package queue

import (
	"github.com/AndrXxX/go-loyalty-service/internal/services/logger"
	"go.uber.org/zap"
)

type worker struct {
}

func (w *worker) Process(jobs <-chan queueJob) {
	for job := range jobs {
		err := job.Execute()
		if err != nil {
			logger.Log.Error("failed to execute runner job", zap.Error(err), zap.Any("job", job))
		}
	}
}
