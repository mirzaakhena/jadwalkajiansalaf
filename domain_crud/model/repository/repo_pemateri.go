package repository

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
)

type SavePemateriRepo interface {
	SavePemateri(ctx context.Context, obj *entity.Pemateri) error
}

type FindAllPemateriFilterRequest struct {
	Page int
	Size int
	// add other field to filter here ...
}

type FindAllPemateriRepo interface {
	FindAllPemateri(ctx context.Context, req FindAllPemateriFilterRequest) ([]*entity.Pemateri, int64, error)
}

type DeletePemateriRepo interface {
	DeletePemateri(ctx context.Context, pemateriID vo.PemateriID) error
}

type FindOnePemateriByIDRepo interface {
	FindOnePemateriByID(ctx context.Context, pemateriID vo.PemateriID) (*entity.Pemateri, error)
}
