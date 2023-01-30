package getonepenyelenggara

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOnePenyelenggaraInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOnePenyelenggaraInteractor{
		outport: outputPort,
	}
}

func (r *GetOnePenyelenggaraInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraObj, err := r.outport.FindOnePenyelenggaraByID(ctx, req.PenyelenggaraID)
	if err != nil {
		return nil, err
	}

	if penyelenggaraObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.Penyelenggara = penyelenggaraObj

	return res, nil
}
