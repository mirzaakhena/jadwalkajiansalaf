package runpemateridelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunPemateriDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPemateriDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunPemateriDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = r.outport.DeletePemateri(ctx, req.PemateriID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
