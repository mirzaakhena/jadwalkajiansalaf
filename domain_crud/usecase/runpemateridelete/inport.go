package runpemateridelete

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	PemateriID vo.PemateriID
}

type InportResponse struct {
}
