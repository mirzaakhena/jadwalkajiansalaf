package runkategorikajianupdate

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneKategoriKajianByIDRepo
	repository.SaveKategoriKajianRepo
}
