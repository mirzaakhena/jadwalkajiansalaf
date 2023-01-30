package runalamatkajiancreate

import (
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.AlamatKajianCreateRequest
}

type InportResponse struct {
	AlamatKajianID vo.AlamatKajianID
}
