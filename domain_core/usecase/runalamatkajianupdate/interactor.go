package runalamatkajianupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunAlamatKajianUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAlamatKajianUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunAlamatKajianUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = alamatKajianObj.Update(req.AlamatKajianUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveAlamatKajian(ctx, alamatKajianObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
