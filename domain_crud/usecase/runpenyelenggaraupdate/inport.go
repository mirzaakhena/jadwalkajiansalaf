package runpenyelenggaraupdate

import (
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.PenyelenggaraUpdateRequest
	PenyelenggaraID vo.PenyelenggaraID
}

type InportResponse struct {
}
