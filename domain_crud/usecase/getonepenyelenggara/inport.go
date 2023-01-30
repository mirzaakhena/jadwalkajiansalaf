package getonepenyelenggara

import (
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	PenyelenggaraID vo.PenyelenggaraID
}

type InportResponse struct {
	Penyelenggara *entity.Penyelenggara
}
