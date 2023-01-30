package withgormdb

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

func (r *gateway) FindAllPenyelenggara(ctx context.Context, req repository.FindAllPenyelenggaraFilterRequest) ([]*entity.Penyelenggara, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOnePenyelenggaraByID(ctx context.Context, penyelenggaraID vo.PenyelenggaraID) (*entity.Penyelenggara, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SavePenyelenggara(ctx context.Context, obj *entity.Penyelenggara) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeletePenyelenggara(ctx context.Context, penyelenggaraID vo.PenyelenggaraID) error {
	r.log.Info(ctx, "called")

	return nil
}
