package runalamatkajianupdate

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneAlamatKajianByIDRepo
	repository.SaveAlamatKajianRepo
}
