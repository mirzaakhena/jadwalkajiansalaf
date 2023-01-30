package getonekajian

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOneKajianByIDRepo
}
