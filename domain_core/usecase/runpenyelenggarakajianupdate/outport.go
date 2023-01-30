package runpenyelenggarakajianupdate

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOnePenyelenggaraKajianByIDRepo
	repository.SavePenyelenggaraKajianRepo
}
