package runkategoriupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunKategoriUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKategoriUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunKategoriUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriObj, err := r.outport.FindOneKategoriByID(ctx, req.KategoriID)
	if err != nil {
		return nil, err
	}
	if kategoriObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = kategoriObj.Update(req.KategoriUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveKategori(ctx, kategoriObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
