package runpenyelenggarakajiandelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunPenyelenggaraKajianDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPenyelenggaraKajianDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunPenyelenggaraKajianDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraKajianObj, err := r.outport.FindOnePenyelenggaraKajianByID(ctx, req.PenyelenggaraKajianID)
	if err != nil {
		return nil, err
	}
	if penyelenggaraKajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = r.outport.DeletePenyelenggaraKajian(ctx, req.PenyelenggaraKajianID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
