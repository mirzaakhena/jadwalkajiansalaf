package repository

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
)

type SavePenyelenggaraKajianRepo interface {
	SavePenyelenggaraKajian(ctx context.Context, obj *entity.PenyelenggaraKajian) error
}

type FindAllPenyelenggaraKajianFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllPenyelenggaraKajianRepo interface {
	FindAllPenyelenggaraKajian(ctx context.Context, req FindAllPenyelenggaraKajianFilterRequest) ([]*entity.PenyelenggaraKajian, int64, error)
}

type DeletePenyelenggaraKajianRepo interface {
	DeletePenyelenggaraKajian(ctx context.Context, penyelenggaraKajianID vo.PenyelenggaraKajianID) error
}

type FindOnePenyelenggaraKajianByIDRepo interface {
	FindOnePenyelenggaraKajianByID(ctx context.Context, penyelenggaraKajianID vo.PenyelenggaraKajianID) (*entity.PenyelenggaraKajian, error)
}
