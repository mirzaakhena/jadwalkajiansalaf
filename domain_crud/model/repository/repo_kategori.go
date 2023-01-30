package repository

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

type SaveKategoriRepo interface {
	SaveKategori(ctx context.Context, obj *entity.Kategori) error
}

type FindAllKategoriFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllKategoriRepo interface {
	FindAllKategori(ctx context.Context, req FindAllKategoriFilterRequest) ([]*entity.Kategori, int64, error)
}

type DeleteKategoriRepo interface {
	DeleteKategori(ctx context.Context, kategoriID vo.KategoriID) error
}

type FindOneKategoriByIDRepo interface {
	FindOneKategoriByID(ctx context.Context, kategoriID vo.KategoriID) (*entity.Kategori, error)
}
