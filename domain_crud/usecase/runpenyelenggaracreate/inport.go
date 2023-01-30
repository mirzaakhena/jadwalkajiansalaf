package runpenyelenggaracreate

import (
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.PenyelenggaraCreateRequest
}

type InportResponse struct {
	PenyelenggaraID vo.PenyelenggaraID
}
