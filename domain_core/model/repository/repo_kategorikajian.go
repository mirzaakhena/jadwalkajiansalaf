package repository

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
)

type SaveKategoriKajianRepo interface {
	SaveKategoriKajian(ctx context.Context, obj *entity.KategoriKajian) error
}

type FindAllKategoriKajianFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllKategoriKajianRepo interface {
	FindAllKategoriKajian(ctx context.Context, req FindAllKategoriKajianFilterRequest) ([]*entity.KategoriKajian, int64, error)
}

type DeleteKategoriKajianRepo interface {
	DeleteKategoriKajian(ctx context.Context, kategoriKajianID vo.KategoriKajianID) error
}

type FindOneKategoriKajianByIDRepo interface {
	FindOneKategoriKajianByID(ctx context.Context, kategoriKajianID vo.KategoriKajianID) (*entity.KategoriKajian, error)
}
