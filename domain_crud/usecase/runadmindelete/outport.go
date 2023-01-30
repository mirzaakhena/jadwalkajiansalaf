package runadmindelete

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOneAdminByIDRepo
	repository.DeleteAdminRepo
}
