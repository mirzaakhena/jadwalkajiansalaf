package getonealamat

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOneAlamatInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOneAlamatInteractor{
		outport: outputPort,
	}
}

func (r *GetOneAlamatInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatObj, err := r.outport.FindOneAlamatByID(ctx, req.AlamatID)
	if err != nil {
		return nil, err
	}

	if alamatObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.Alamat = alamatObj

	return res, nil
}
