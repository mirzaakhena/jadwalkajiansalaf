package runalamatupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunAlamatUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAlamatUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunAlamatUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = alamatObj.Update(req.LokasiUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveAlamat(ctx, alamatObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
