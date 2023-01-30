package getonealamatkajian

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOneAlamatKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOneAlamatKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetOneAlamatKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatKajianObj, err := r.outport.FindOneAlamatKajianByID(ctx, req.AlamatKajianID)
	if err != nil {
		return nil, err
	}

	if alamatKajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.AlamatKajian = alamatKajianObj

	return res, nil
}
