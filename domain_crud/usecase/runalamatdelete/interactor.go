package runalamatdelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunAlamatDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAlamatDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunAlamatDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatObj, err := r.outport.FindOneAlamatByID(ctx, req.AlamatID)
	if err != nil {
		return nil, err
	}
	if alamatObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = r.outport.DeleteAlamat(ctx, req.AlamatID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
