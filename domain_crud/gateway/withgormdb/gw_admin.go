package withgormdb

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

func (r *gateway) FindAllAdmin(ctx context.Context, req repository.FindAllAdminFilterRequest) ([]*entity.Admin, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneAdminByID(ctx context.Context, adminID vo.AdminID) (*entity.Admin, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveAdmin(ctx context.Context, obj *entity.Admin) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteAdmin(ctx context.Context, adminID vo.AdminID) error {
	r.log.Info(ctx, "called")

	return nil
}
