package withgormdb2

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/config"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&entity.Admin{})
	if err != nil {
		panic("cannot create schema")
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		db:      db,
	}
}

func (r *gateway) FindAllAdmin(ctx context.Context, req repository.FindAllAdminFilterRequest) ([]*entity.Admin, int64, error) {
	r.log.Info(ctx, "called")

	var result []*entity.Admin
	err := r.db.Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return nil, 0, nil
}

func (r *gateway) FindAllAlamat(ctx context.Context, req repository.FindAllAlamatFilterRequest) ([]*entity.Lokasi, int64, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, 0, nil
}

func (r *gateway) FindAllKajian(ctx context.Context, req repository.FindAllKajianFilterRequest) ([]*entity.Kajian, int64, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, 0, nil
}

func (r *gateway) FindAllKategori(ctx context.Context, req repository.FindAllKategoriFilterRequest) ([]*entity.Kategori, int64, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, 0, nil
}

func (r *gateway) FindAllPemateri(ctx context.Context, req repository.FindAllPemateriFilterRequest) ([]*entity.Pemateri, int64, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, 0, nil
}

func (r *gateway) FindAllPenyelenggara(ctx context.Context, req repository.FindAllPenyelenggaraFilterRequest) ([]*entity.Penyelenggara, int64, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, 0, nil
}

func (r *gateway) FindOneAdminByID(ctx context.Context, adminID vo.AdminID) (*entity.Admin, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, nil
}

func (r *gateway) FindOneAlamatByID(ctx context.Context, alamatID vo.LokasiID) (*entity.Lokasi, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, nil
}

func (r *gateway) FindOneKajianByID(ctx context.Context, kajianID vo.KajianID) (*entity.Kajian, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, nil
}

func (r *gateway) FindOneKategoriByID(ctx context.Context, kategoriID vo.KategoriID) (*entity.Kategori, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, nil
}

func (r *gateway) FindOnePemateriByID(ctx context.Context, pemateriID vo.PemateriID) (*entity.Pemateri, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, nil
}

func (r *gateway) FindOnePenyelenggaraByID(ctx context.Context, penyelenggaraID vo.PenyelenggaraID) (*entity.Penyelenggara, error) {
	r.log.Info(ctx, "called")

	//var result []*entity.Person
	//err := r.db.Find(&result).Error
	//if err != nil {
	//    return nil, err
	//}

	return nil, nil
}

func (r *gateway) SaveAdmin(ctx context.Context, obj *entity.Admin) error {
	r.log.Info(ctx, "called")

	err := r.db.Create(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) DeleteAdmin(ctx context.Context, adminID vo.AdminID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveAlamat(ctx context.Context, obj *entity.Lokasi) error {
	r.log.Info(ctx, "called")

	// err := r.db.Create(obj).Error
	// if err != nil {
	//     return err
	// }

	return nil
}

func (r *gateway) DeleteAlamat(ctx context.Context, alamatID vo.LokasiID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveKajian(ctx context.Context, obj *entity.Kajian) error {
	r.log.Info(ctx, "called")

	// err := r.db.Create(obj).Error
	// if err != nil {
	//     return err
	// }

	return nil
}

func (r *gateway) DeleteKajian(ctx context.Context, kajianID vo.KajianID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveKategori(ctx context.Context, obj *entity.Kategori) error {
	r.log.Info(ctx, "called")

	// err := r.db.Create(obj).Error
	// if err != nil {
	//     return err
	// }

	return nil
}

func (r *gateway) DeleteKategori(ctx context.Context, kategoriID vo.KategoriID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SavePemateri(ctx context.Context, obj *entity.Pemateri) error {
	r.log.Info(ctx, "called")

	// err := r.db.Create(obj).Error
	// if err != nil {
	//     return err
	// }

	return nil
}

func (r *gateway) DeletePemateri(ctx context.Context, pemateriID vo.PemateriID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SavePenyelenggara(ctx context.Context, obj *entity.Penyelenggara) error {
	r.log.Info(ctx, "called")

	// err := r.db.Create(obj).Error
	// if err != nil {
	//     return err
	// }

	return nil
}

func (r *gateway) DeletePenyelenggara(ctx context.Context, penyelenggaraID vo.PenyelenggaraID) error {
	r.log.Info(ctx, "called")

	return nil
}
