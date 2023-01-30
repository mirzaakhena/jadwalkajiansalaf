package runkajianupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunKajianUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKajianUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunKajianUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kajianObj, err := r.outport.FindOneKajianByID(ctx, req.KajianID)
	if err != nil {
		return nil, err
	}
	if kajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = kajianObj.Update(req.KajianUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveKajian(ctx, kajianObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
