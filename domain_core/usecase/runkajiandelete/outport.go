package runkajiandelete

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneKajianByIDRepo
	repository.DeleteKajianRepo
}
