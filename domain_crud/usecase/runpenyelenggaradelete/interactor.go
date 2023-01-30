package runpenyelenggaradelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunPenyelenggaraDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPenyelenggaraDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunPenyelenggaraDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraObj, err := r.outport.FindOnePenyelenggaraByID(ctx, req.PenyelenggaraID)
	if err != nil {
		return nil, err
	}
	if penyelenggaraObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = r.outport.DeletePenyelenggara(ctx, req.PenyelenggaraID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
