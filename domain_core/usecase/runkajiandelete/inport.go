package runkajiandelete

import (
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	KajianID vo.KajianID
}

type InportResponse struct {
}
