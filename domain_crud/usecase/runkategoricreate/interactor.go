package runkategoricreate

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunKategoriCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKategoriCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunKategoriCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriObj, err := entity.NewKategori(req.KategoriCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveKategori(ctx, kategoriObj)
	if err != nil {
		return nil, err
	}

	res.KategoriID = kategoriObj.ID

	return res, nil
}
