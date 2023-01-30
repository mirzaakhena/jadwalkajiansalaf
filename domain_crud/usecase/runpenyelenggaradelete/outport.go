package runpenyelenggaradelete

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOnePenyelenggaraByIDRepo
	repository.DeletePenyelenggaraRepo
}
