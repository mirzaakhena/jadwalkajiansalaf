package runkategorikajiandelete

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneKategoriKajianByIDRepo
	repository.DeleteKategoriKajianRepo
}
