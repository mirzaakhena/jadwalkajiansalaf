package runkajiandelete

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	KajianID vo.KajianID
}

type InportResponse struct {
}
