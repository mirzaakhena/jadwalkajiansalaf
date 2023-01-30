package getonekajian

import (
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	KajianID vo.KajianID
}

type InportResponse struct {
	Kajian *entity.Kajian
}
