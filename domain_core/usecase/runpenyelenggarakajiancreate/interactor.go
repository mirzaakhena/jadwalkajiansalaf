package runpenyelenggarakajiancreate

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunPenyelenggaraKajianCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPenyelenggaraKajianCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunPenyelenggaraKajianCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraKajianObj, err := entity.NewPenyelenggaraKajian(req.PenyelenggaraKajianCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SavePenyelenggaraKajian(ctx, penyelenggaraKajianObj)
	if err != nil {
		return nil, err
	}

	res.PenyelenggaraKajianID = penyelenggaraKajianObj.ID

	return res, nil
}
