package runpemateriupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunPemateriUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPemateriUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunPemateriUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	pemateriObj, err := r.outport.FindOnePemateriByID(ctx, req.PemateriID)
	if err != nil {
		return nil, err
	}
	if pemateriObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = pemateriObj.Update(req.PemateriUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SavePemateri(ctx, pemateriObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
