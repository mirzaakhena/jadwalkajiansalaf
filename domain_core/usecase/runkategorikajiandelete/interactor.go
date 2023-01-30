package runkategorikajiandelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunKategoriKajianDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKategoriKajianDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunKategoriKajianDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = r.outport.DeleteKategoriKajian(ctx, req.KategoriKajianID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
