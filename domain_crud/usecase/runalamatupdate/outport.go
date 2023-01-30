package runalamatupdate

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOneAlamatByIDRepo
	repository.SaveAlamatRepo
}
