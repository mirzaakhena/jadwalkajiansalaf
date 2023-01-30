package getallkategorikajian

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllKategoriKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllKategoriKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetAllKategoriKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriKajianObj, count, err := r.outport.FindAllKategoriKajian(ctx, req.FindAllKategoriKajianFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(kategoriKajianObj)

	return res, nil
}
