package runalamatkajiancreate

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunAlamatKajianCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAlamatKajianCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunAlamatKajianCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatKajianObj, err := entity.NewAlamatKajian(req.AlamatKajianCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveAlamatKajian(ctx, alamatKajianObj)
	if err != nil {
		return nil, err
	}

	res.AlamatKajianID = alamatKajianObj.ID

	return res, nil
}
