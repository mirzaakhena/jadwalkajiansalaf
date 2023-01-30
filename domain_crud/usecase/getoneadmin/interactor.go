package getoneadmin

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type GetOneAdminInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetOneAdminInteractor{
		outport: outputPort,
	}
}

func (r *GetOneAdminInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	adminObj, err := r.outport.FindOneAdminByID(ctx, req.AdminID)
	if err != nil {
		return nil, err
	}

	if adminObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.Admin = adminObj

	return res, nil
}
