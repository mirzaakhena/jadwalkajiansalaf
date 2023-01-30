package runadminupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type RunAdminUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAdminUpdateInteractor{
		outport: outputPort,
	}
}

func (r *RunAdminUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = adminObj.Update(req.AdminUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveAdmin(ctx, adminObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
