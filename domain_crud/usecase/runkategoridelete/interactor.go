package runkategoridelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunKategoriDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKategoriDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunKategoriDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = r.outport.DeleteKategori(ctx, req.KategoriID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
