package getonepemateri

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOnePemateriInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOnePemateriInteractor{
		outport: outputPort,
	}
}

func (r *GetOnePemateriInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	pemateriObj, err := r.outport.FindOnePemateriByID(ctx, req.PemateriID)
	if err != nil {
		return nil, err
	}

	if pemateriObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.Pemateri = pemateriObj

	return res, nil
}
