package withgormdb

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

func (r *gateway) FindAllKajian(ctx context.Context, req repository.FindAllKajianFilterRequest) ([]*entity.Kajian, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneKajianByID(ctx context.Context, kajianID vo.KajianID) (*entity.Kajian, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveKajian(ctx context.Context, obj *entity.Kajian) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteKajian(ctx context.Context, kajianID vo.KajianID) error {
	r.log.Info(ctx, "called")

	return nil
}
