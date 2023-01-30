package repository

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

type SaveKajianRepo interface {
	SaveKajian(ctx context.Context, obj *entity.Kajian) error
}

type FindAllKajianFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllKajianRepo interface {
	FindAllKajian(ctx context.Context, req FindAllKajianFilterRequest) ([]*entity.Kajian, int64, error)
}

type DeleteKajianRepo interface {
	DeleteKajian(ctx context.Context, kajianID vo.KajianID) error
}

type FindOneKajianByIDRepo interface {
	FindOneKajianByID(ctx context.Context, kajianID vo.KajianID) (*entity.Kajian, error)
}
