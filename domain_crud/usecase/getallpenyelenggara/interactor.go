package getallpenyelenggara

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllPenyelenggaraInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllPenyelenggaraInteractor{
		outport: outputPort,
	}
}

func (r *GetAllPenyelenggaraInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	penyelenggaraObj, count, err := r.outport.FindAllPenyelenggara(ctx, req.FindAllPenyelenggaraFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(penyelenggaraObj)

	return res, nil
}
