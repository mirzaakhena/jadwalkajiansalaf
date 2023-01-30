package runkajiancreate

import (
	"jadwalkajiansalaf/domain_core/model/repository"
)

type Outport interface {
	repository.SaveKajianRepo
}
