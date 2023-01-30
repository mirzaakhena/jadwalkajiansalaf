package getallalamatkajian

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindAllAlamatKajianRepo
}
