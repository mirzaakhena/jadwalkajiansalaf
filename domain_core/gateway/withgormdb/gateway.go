package withgormdb

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/repository"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/config"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
)

type gateway struct {
	appData gogen.ApplicationData
	config  *config.Config
	log     logger.Logger
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) FindAllAdmin(ctx context.Context, req repository.FindAllAdminFilterRequest) ([]*entity.Admin, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllAlamatKajian(ctx context.Context, req repository.FindAllAlamatKajianFilterRequest) ([]*entity.AlamatKajian, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllKajian(ctx context.Context, req repository.FindAllKajianFilterRequest) ([]*entity.Kajian, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllKategoriKajian(ctx context.Context, req repository.FindAllKategoriKajianFilterRequest) ([]*entity.KategoriKajian, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllPemateri(ctx context.Context, req repository.FindAllPemateriFilterRequest) ([]*entity.Pemateri, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllPenyelenggaraKajian(ctx context.Context, req repository.FindAllPenyelenggaraKajianFilterRequest) ([]*entity.PenyelenggaraKajian, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneAdminByID(ctx context.Context, adminID vo.AdminID) (*entity.Admin, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindOneAlamatKajianByID(ctx context.Context, alamatKajianID vo.AlamatKajianID) (*entity.AlamatKajian, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindOneKajianByID(ctx context.Context, kajianID vo.KajianID) (*entity.Kajian, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindOneKategoriKajianByID(ctx context.Context, kategoriKajianID vo.KategoriKajianID) (*entity.KategoriKajian, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindOnePemateriByID(ctx context.Context, pemateriID vo.PemateriID) (*entity.Pemateri, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindOnePenyelenggaraKajianByID(ctx context.Context, penyelenggaraKajianID vo.PenyelenggaraKajianID) (*entity.PenyelenggaraKajian, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveAdmin(ctx context.Context, obj *entity.Admin) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteAdmin(ctx context.Context, adminID vo.AdminID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveAlamatKajian(ctx context.Context, obj *entity.AlamatKajian) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteAlamatKajian(ctx context.Context, alamatKajianID vo.AlamatKajianID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveKajian(ctx context.Context, obj *entity.Kajian) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteKajian(ctx context.Context, kajianID vo.KajianID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveKategoriKajian(ctx context.Context, obj *entity.KategoriKajian) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteKategoriKajian(ctx context.Context, kategoriKajianID vo.KategoriKajianID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SavePemateri(ctx context.Context, obj *entity.Pemateri) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeletePemateri(ctx context.Context, pemateriID vo.PemateriID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SavePenyelenggaraKajian(ctx context.Context, obj *entity.PenyelenggaraKajian) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeletePenyelenggaraKajian(ctx context.Context, penyelenggaraKajianID vo.PenyelenggaraKajianID) error {
	r.log.Info(ctx, "called")

	return nil
}
