package runpenyelenggaracreate

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunPenyelenggaraCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPenyelenggaraCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunPenyelenggaraCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraObj, err := entity.NewPenyelenggara(req.PenyelenggaraCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SavePenyelenggara(ctx, penyelenggaraObj)
	if err != nil {
		return nil, err
	}

	res.PenyelenggaraID = penyelenggaraObj.ID

	return res, nil
}
