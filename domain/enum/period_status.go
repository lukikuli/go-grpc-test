package enum

import (
	"errors"
	"strings"
)

type PeriodStatus struct {
	value string
}

const (
	PERIOD_STATUS_CURRENT = "CURRENT"
	PERIOD_STATUS_ENDED   = "ENDED"
)

var validPeriodStatus = map[string]struct{}{
	PERIOD_STATUS_CURRENT: {},
	PERIOD_STATUS_ENDED:   {},
}

func NewPeriodStatusByString(value string) (PeriodStatus, error) {
	if _, exists := validPeriodStatus[strings.ToUpper(value)]; exists {
		return PeriodStatus{value: strings.ToUpper(value)}, nil
	}

	return PeriodStatus{}, errors.New("period status not found")
}

func (ps PeriodStatus) GetStringValue() string {
	return ps.value
}

func (ps PeriodStatus) IsCurrent() bool {
	return ps.value == PERIOD_STATUS_CURRENT
}

func (ps PeriodStatus) IsEnded() bool {
	return ps.value == PERIOD_STATUS_ENDED
}
