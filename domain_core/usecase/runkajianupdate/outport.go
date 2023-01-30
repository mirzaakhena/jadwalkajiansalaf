package runkajianupdate

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneKajianByIDRepo
	repository.SaveKajianRepo
}
