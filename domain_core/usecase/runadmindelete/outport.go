package runadmindelete

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.FindOneAdminByIDRepo
	repository.DeleteAdminRepo
}
