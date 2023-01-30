package runalamatcreate

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunAlamatCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAlamatCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunAlamatCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatObj, err := entity.NewLokasi(req.LokasiCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveAlamat(ctx, alamatObj)
	if err != nil {
		return nil, err
	}

	res.AlamatID = alamatObj.ID

	return res, nil
}
