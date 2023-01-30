package getonekajian

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneKajianByIDRepo
}
