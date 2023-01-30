package runpematericreate

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunPemateriCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunPemateriCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunPemateriCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	pemateriObj, err := entity.NewPemateri(req.PemateriCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SavePemateri(ctx, pemateriObj)
	if err != nil {
		return nil, err
	}

	res.PemateriID = pemateriObj.ID

	return res, nil
}
