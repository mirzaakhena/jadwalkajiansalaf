package repository

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

type SavePenyelenggaraRepo interface {
	SavePenyelenggara(ctx context.Context, obj *entity.Penyelenggara) error
}

type FindAllPenyelenggaraFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllPenyelenggaraRepo interface {
	FindAllPenyelenggara(ctx context.Context, req FindAllPenyelenggaraFilterRequest) ([]*entity.Penyelenggara, int64, error)
}

type DeletePenyelenggaraRepo interface {
	DeletePenyelenggara(ctx context.Context, penyelenggaraID vo.PenyelenggaraID) error
}

type FindOnePenyelenggaraByIDRepo interface {
	FindOnePenyelenggaraByID(ctx context.Context, penyelenggaraID vo.PenyelenggaraID) (*entity.Penyelenggara, error)
}
