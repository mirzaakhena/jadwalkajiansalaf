package getallkajian

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetAllKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kajianObj, count, err := r.outport.FindAllKajian(ctx, req.FindAllKajianFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(kajianObj)

	return res, nil
}
