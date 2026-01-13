package entity

import (
	"errors"
	"go-grpc/domain/enum"
	"time"
)

type Period struct {
	periodId     int
	periodName   string
	periodStart  time.Time
	periodEnd    time.Time
	expTreshold  int
	periodStatus enum.PeriodStatus
}

type PeriodParam struct {
	PeriodId    int
	PeriodName  string
	PeriodStart time.Time
	PeriodEnd   time.Time
	ExpTreshold int
}

func NewPeriod(param PeriodParam) (*Period, error) {
	if param.PeriodStart.After(param.PeriodEnd) {
		return nil, errors.New("invalid period start and period end")
	}

	period := &Period{
		periodId:    param.PeriodId,
		periodName:  param.PeriodName,
		periodStart: param.PeriodStart,
		periodEnd:   param.PeriodEnd,
		expTreshold: param.ExpTreshold,
	}

	periodStatus, _ := enum.NewPeriodStatusByString(enum.PERIOD_STATUS_ENDED)
	if period.IsActive() {
		periodStatus, _ = enum.NewPeriodStatusByString(enum.PERIOD_STATUS_CURRENT)
	}
	period.periodStatus = periodStatus

	return period, nil
}

func (p *Period) GetPeriodId() int {
	return p.periodId
}

func (p *Period) GetPeriodName() string {
	return p.periodName
}

func (p *Period) GetPeriodStart() time.Time {
	return p.periodStart
}

func (p *Period) GetPeriodEnd() time.Time {
	return p.periodEnd
}

func (p *Period) GetPeriodExpTreshold() int {
	return p.expTreshold
}

func (p *Period) GetPeriodStatus() enum.PeriodStatus {
	return p.periodStatus
}

func (p *Period) IsExpired() bool {
	return time.Now().After(p.periodEnd)
}

func (p *Period) IsUpcoming() bool {
	return time.Now().Before(p.periodStart)
}

func (p *Period) IsActive() bool {
	return !time.Now().Before(p.periodStart) && !time.Now().After(p.periodEnd)
}
