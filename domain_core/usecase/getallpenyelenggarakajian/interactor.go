package getallpenyelenggarakajian

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllPenyelenggaraKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllPenyelenggaraKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetAllPenyelenggaraKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraKajianObj, count, err := r.outport.FindAllPenyelenggaraKajian(ctx, req.FindAllPenyelenggaraKajianFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(penyelenggaraKajianObj)

	return res, nil
}
