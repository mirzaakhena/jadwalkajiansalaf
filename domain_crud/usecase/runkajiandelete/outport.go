package runkajiandelete

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOneKajianByIDRepo
	repository.DeleteKajianRepo
}
