package runpenyelenggarakajianupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunPenyelenggaraKajianUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPenyelenggaraKajianUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunPenyelenggaraKajianUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = penyelenggaraKajianObj.Update(req.PenyelenggaraKajianUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SavePenyelenggaraKajian(ctx, penyelenggaraKajianObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
