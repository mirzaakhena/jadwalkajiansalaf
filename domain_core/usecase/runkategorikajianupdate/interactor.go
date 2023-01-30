package runkategorikajianupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunKategoriKajianUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKategoriKajianUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunKategoriKajianUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriKajianObj, err := r.outport.FindOneKategoriKajianByID(ctx, req.KategoriKajianID)
	if err != nil {
		return nil, err
	}
	if kategoriKajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = kategoriKajianObj.Update(req.KategoriKajianUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveKategoriKajian(ctx, kategoriKajianObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
