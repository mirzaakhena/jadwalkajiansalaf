package getalladmin

import (
	"context"
	"jadwalkajiansalaf/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type GetAllAdminInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &GetAllAdminInteractor{
		outport: outputPort,
	}
}

func (r *GetAllAdminInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	adminObj, count, err := r.outport.FindAllAdmin(ctx, req.FindAllAdminFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(adminObj)

	return res, nil
}
