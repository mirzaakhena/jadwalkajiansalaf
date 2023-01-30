package runkategoriupdate

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOneKategoriByIDRepo
	repository.SaveKategoriRepo
}
