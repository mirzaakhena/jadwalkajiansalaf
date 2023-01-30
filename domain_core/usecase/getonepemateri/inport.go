package getonepemateri

import (
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	PemateriID vo.PemateriID
}

type InportResponse struct {
	Pemateri *entity.Pemateri
}
