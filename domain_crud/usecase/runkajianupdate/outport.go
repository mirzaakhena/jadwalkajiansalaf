package runkajianupdate

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOneKajianByIDRepo
	repository.SaveKajianRepo
}
