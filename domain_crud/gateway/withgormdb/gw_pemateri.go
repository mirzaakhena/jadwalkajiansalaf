package withgormdb

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

func (r *gateway) FindAllPemateri(ctx context.Context, req repository.FindAllPemateriFilterRequest) ([]*entity.Pemateri, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOnePemateriByID(ctx context.Context, pemateriID vo.PemateriID) (*entity.Pemateri, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SavePemateri(ctx context.Context, obj *entity.Pemateri) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeletePemateri(ctx context.Context, pemateriID vo.PemateriID) error {
	r.log.Info(ctx, "called")

	return nil
}
