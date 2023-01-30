package runkajiancreate

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunKajianCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunKajianCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunKajianCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kajianObj, err := entity.NewKajian(req.KajianCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveKajian(ctx, kajianObj)
	if err != nil {
		return nil, err
	}

	res.KajianID = kajianObj.ID

	return res, nil
}
