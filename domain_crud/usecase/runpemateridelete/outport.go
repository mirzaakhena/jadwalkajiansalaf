package runpemateridelete

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
)

type Outport interface {
	repository.FindOnePemateriByIDRepo
	repository.DeletePemateriRepo
}
