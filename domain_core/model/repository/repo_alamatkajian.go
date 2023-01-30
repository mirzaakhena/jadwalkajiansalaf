package repository

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
)

type SaveAlamatKajianRepo interface {
	SaveAlamatKajian(ctx context.Context, obj *entity.AlamatKajian) error
}

type FindAllAlamatKajianFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllAlamatKajianRepo interface {
	FindAllAlamatKajian(ctx context.Context, req FindAllAlamatKajianFilterRequest) ([]*entity.AlamatKajian, int64, error)
}

type DeleteAlamatKajianRepo interface {
	DeleteAlamatKajian(ctx context.Context, alamatKajianID vo.AlamatKajianID) error
}

type FindOneAlamatKajianByIDRepo interface {
	FindOneAlamatKajianByID(ctx context.Context, alamatKajianID vo.AlamatKajianID) (*entity.AlamatKajian, error)
}
