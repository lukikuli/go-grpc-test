package domain

import (
	"context"
	"go-grpc/domain/entity"
	"time"
)

type PeriodRepository interface {
	GetCurrentActivePeriod(ctx context.Context) (*entity.Period, error)
	GetLastPeriod(ctx context.Context) (*entity.Period, error)
	GetMasterPeriods(ctx context.Context) ([]*entity.Period, error)
	GetPeriodByID(ctx context.Context, id int) (*entity.Period, error)
	GetPeriodByRange(ctx context.Context, startDate, endDate time.Time) (*entity.Period, error)
}
