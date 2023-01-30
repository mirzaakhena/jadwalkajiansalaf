package getallalamat

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllAlamatInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllAlamatInteractor{
		outport: outputPort,
	}
}

func (r *GetAllAlamatInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatObj, count, err := r.outport.FindAllAlamat(ctx, req.FindAllAlamatFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(alamatObj)

	return res, nil
}
