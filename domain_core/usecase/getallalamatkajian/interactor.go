package getallalamatkajian

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllAlamatKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllAlamatKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetAllAlamatKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	alamatKajianObj, count, err := r.outport.FindAllAlamatKajian(ctx, req.FindAllAlamatKajianFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(alamatKajianObj)

	return res, nil
}
