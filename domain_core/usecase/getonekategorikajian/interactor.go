package getonekategorikajian

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOneKategoriKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOneKategoriKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetOneKategoriKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kategoriKajianObj, err := r.outport.FindOneKategoriKajianByID(ctx, req.KategoriKajianID)
	if err != nil {
		return nil, err
	}

	if kategoriKajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.KategoriKajian = kategoriKajianObj

	return res, nil
}
