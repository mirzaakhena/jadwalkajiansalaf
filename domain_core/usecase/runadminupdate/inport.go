package runadminupdate

import (
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.AdminUpdateRequest
	AdminID vo.AdminID
}

type InportResponse struct {
}
