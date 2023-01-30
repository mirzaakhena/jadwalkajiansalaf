package repository

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

type SaveAlamatRepo interface {
	SaveAlamat(ctx context.Context, obj *entity.Lokasi) error
}

type FindAllAlamatFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllAlamatRepo interface {
	FindAllAlamat(ctx context.Context, req FindAllAlamatFilterRequest) ([]*entity.Lokasi, int64, error)
}

type DeleteAlamatRepo interface {
	DeleteAlamat(ctx context.Context, alamatID vo.LokasiID) error
}

type FindOneAlamatByIDRepo interface {
	FindOneAlamatByID(ctx context.Context, alamatID vo.LokasiID) (*entity.Lokasi, error)
}
