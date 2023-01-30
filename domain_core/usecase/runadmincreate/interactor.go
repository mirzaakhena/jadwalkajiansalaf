package runadmincreate

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type RunAdminCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &RunAdminCreateInteractor{
		outport: outputPort,
	}
}

func (r *RunAdminCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	adminObj, err := entity.NewAdmin(req.AdminCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveAdmin(ctx, adminObj)
	if err != nil {
		return nil, err
	}

	res.AdminID = adminObj.ID

	return res, nil
}
