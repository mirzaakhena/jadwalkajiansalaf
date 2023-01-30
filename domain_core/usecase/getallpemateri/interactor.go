package getallpemateri

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllPemateriInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllPemateriInteractor{
		outport: outputPort,
	}
}

func (r *GetAllPemateriInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	pemateriObj, count, err := r.outport.FindAllPemateri(ctx, req.FindAllPemateriFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(pemateriObj)

	return res, nil
}
