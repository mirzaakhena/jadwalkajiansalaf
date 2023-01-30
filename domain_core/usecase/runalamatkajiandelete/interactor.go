package runalamatkajiandelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunAlamatKajianDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAlamatKajianDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunAlamatKajianDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatKajianObj, err := r.outport.FindOneAlamatKajianByID(ctx, req.AlamatKajianID)
	if err != nil {
		return nil, err
	}
	if alamatKajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = r.outport.DeleteAlamatKajian(ctx, req.AlamatKajianID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
