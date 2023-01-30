package getonekategori

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOneKategoriInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOneKategoriInteractor{
		outport: outputPort,
	}
}

func (r *GetOneKategoriInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriObj, err := r.outport.FindOneKategoriByID(ctx, req.KategoriID)
	if err != nil {
		return nil, err
	}

	if kategoriObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.Kategori = kategoriObj

	return res, nil
}
