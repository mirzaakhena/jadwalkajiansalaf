package getonepenyelenggarakajian

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOnePenyelenggaraKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOnePenyelenggaraKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetOnePenyelenggaraKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraKajianObj, err := r.outport.FindOnePenyelenggaraKajianByID(ctx, req.PenyelenggaraKajianID)
	if err != nil {
		return nil, err
	}

	if penyelenggaraKajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.PenyelenggaraKajian = penyelenggaraKajianObj

	return res, nil
}
