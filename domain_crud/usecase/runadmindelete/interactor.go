package runadmindelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunAdminDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAdminDeleteInteractor{
		outport: outputPort,
	}
}

func (r *RunAdminDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	adminObj, err := r.outport.FindOneAdminByID(ctx, req.AdminID)
	if err != nil {
		return nil, err
	}
	if adminObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = r.outport.DeleteAdmin(ctx, req.AdminID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
