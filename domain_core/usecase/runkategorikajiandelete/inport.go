package runkategorikajiandelete

import (
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	KategoriKajianID vo.KategoriKajianID
}

type InportResponse struct {
}
