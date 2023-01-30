package runpenyelenggaraupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunPenyelenggaraUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPenyelenggaraUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunPenyelenggaraUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = penyelenggaraObj.Update(req.PenyelenggaraUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SavePenyelenggara(ctx, penyelenggaraObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
