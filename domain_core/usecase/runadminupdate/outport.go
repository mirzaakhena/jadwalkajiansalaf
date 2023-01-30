package runadminupdate

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneAdminByIDRepo
	repository.SaveAdminRepo
}
