package withgormdb

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

func (r *gateway) FindAllAlamat(ctx context.Context, req repository.FindAllAlamatFilterRequest) ([]*entity.Lokasi, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneAlamatByID(ctx context.Context, alamatID vo.LokasiID) (*entity.Lokasi, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveAlamat(ctx context.Context, obj *entity.Lokasi) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteAlamat(ctx context.Context, alamatID vo.LokasiID) error {
	r.log.Info(ctx, "called")

	return nil
}
