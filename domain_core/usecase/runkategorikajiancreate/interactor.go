package runkategorikajiancreate

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunKategoriKajianCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKategoriKajianCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunKategoriKajianCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriKajianObj, err := entity.NewKategoriKajian(req.KategoriKajianCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveKategoriKajian(ctx, kategoriKajianObj)
	if err != nil {
		return nil, err
	}

	res.KategoriKajianID = kategoriKajianObj.ID

	return res, nil
}
