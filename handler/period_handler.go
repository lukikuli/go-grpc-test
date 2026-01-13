package handler

import (
	"context"
	"go-grpc/protobuf/period"
)

// PeriodHandler is a struct implementing PeriodHandler
type PeriodHandler struct {
	period.UnimplementedPeriodHandlerServer
}

func (p *PeriodHandler) GetPeriodById(ctx context.Context, req *period.GetPeriodRequest) (*period.GetPeriodResponse, error) {
	period := &period.GetPeriodResponse{
		PeriodId:   1,
		PeriodName: "Period Testing 1",
	}
	return period, nil
}
func (p *PeriodHandler) GetPeriodList(ctx context.Context, req *period.Empty) (*period.GetPeriodListResponse, error) {
	periodList := &period.GetPeriodListResponse{
		Data: []*period.GetPeriodResponse{
			&period.GetPeriodResponse{
				PeriodId:   1,
				PeriodName: "Period Testing 1",
			},
			&period.GetPeriodResponse{
				PeriodId:   2,
				PeriodName: "Period Testing 2",
			},
			&period.GetPeriodResponse{
				PeriodId:   3,
				PeriodName: "Period Testing 3",
			},
		},
	}
	return periodList, nil
}
