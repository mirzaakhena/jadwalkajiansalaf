package getonekategori

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOneKategoriByIDRepo
}
