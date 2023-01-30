package getallkategori

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllKategoriInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllKategoriInteractor{
		outport: outputPort,
	}
}

func (r *GetAllKategoriInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriObj, count, err := r.outport.FindAllKategori(ctx, req.FindAllKategoriFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(kategoriObj)

	return res, nil
}
