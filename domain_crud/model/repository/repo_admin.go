package repository

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

type SaveAdminRepo interface {
	SaveAdmin(ctx context.Context, obj *entity.Admin) error
}

type FindAllAdminFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllAdminRepo interface {
	FindAllAdmin(ctx context.Context, req FindAllAdminFilterRequest) ([]*entity.Admin, int64, error)
}

type DeleteAdminRepo interface {
	DeleteAdmin(ctx context.Context, adminID vo.AdminID) error
}

type FindOneAdminByIDRepo interface {
	FindOneAdminByID(ctx context.Context, adminID vo.AdminID) (*entity.Admin, error)
}
