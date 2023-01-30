package withgormdb

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

func (r *gateway) FindAllKategori(ctx context.Context, req repository.FindAllKategoriFilterRequest) ([]*entity.Kategori, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneKategoriByID(ctx context.Context, kategoriID vo.KategoriID) (*entity.Kategori, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveKategori(ctx context.Context, obj *entity.Kategori) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteKategori(ctx context.Context, kategoriID vo.KategoriID) error {
	r.log.Info(ctx, "called")

	return nil
}
