package runkategoridelete

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	KategoriID vo.KategoriID
}

type InportResponse struct {
}
