package getonekajian

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOneKajianInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOneKajianInteractor{
		outport: outputPort,
	}
}

func (r *GetOneKajianInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	kajianObj, err := r.outport.FindOneKajianByID(ctx, req.KajianID)
	if err != nil {
		return nil, err
	}

	if kajianObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.Kajian = kajianObj

	return res, nil
}
